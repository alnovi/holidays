package provider

import (
	"context"
	"log/slog"

	"github.com/go-telegram/bot"

	"github.com/alnovi/holidays/config"
	"github.com/alnovi/holidays/pkg/closer"
	"github.com/alnovi/holidays/pkg/configure"
	"github.com/alnovi/holidays/pkg/database/sqlite"
	"github.com/alnovi/holidays/pkg/logger"
	"github.com/alnovi/holidays/pkg/migrator"
	"github.com/alnovi/holidays/pkg/scheduler"
	"github.com/alnovi/holidays/pkg/utils"
	"github.com/alnovi/holidays/pkg/validator"
)

type Provider struct {
	config    *config.Config
	logger    *slog.Logger
	closer    *closer.Closer
	validator *validator.EchoValidator
	db        *sqlite.Client
	scheduler *scheduler.Scheduler
	telegram  *bot.Bot
}

func New(config *config.Config) *Provider {
	return &Provider{config: config}
}

func (p *Provider) Config() *config.Config {
	if p.config == nil {
		p.config = new(config.Config)
		err := configure.LoadFromEnv(p.config)
		utils.MustMsg(err, "failed to load environment variables config")
		p.config.Normalize()
	}
	return p.config
}

func (p *Provider) Logger() *slog.Logger {
	if p.logger == nil {
		p.logger = logger.New(
			logger.WithFormat(p.Config().Logger.Format),
			logger.WithLevel(p.Config().Logger.Level),
		)
	}
	return p.logger
}

func (p *Provider) LoggerMod(mod string) *slog.Logger {
	if mod == "" {
		return p.Logger()
	}
	return p.Logger().With("module", mod)
}

func (p *Provider) Closer() *closer.Closer {
	if p.closer == nil {
		p.closer = closer.New(p.Config().App.Shutdown)
	}
	return p.closer
}

func (p *Provider) Validator() *validator.EchoValidator {
	if p.validator == nil {
		p.validator = validator.NewEchoValidator()
	}
	return p.validator
}

func (p *Provider) DB() *sqlite.Client {
	if p.db == nil {
		var err error
		p.db, err = sqlite.NewClient(p.Config().Database.Database)
		utils.MustMsg(err, "failed to open database connection")
		utils.MustMsg(err, "fail init db")
		utils.MustMsg(p.db.Ping(context.Background()), "fail ping db")
	}
	return p.db
}

func (p *Provider) MigrationUp() {
	ctx := context.WithValue(context.Background(), migrator.ConfigKey, p.Config())
	log := migrator.NewGooseLogger(p.LoggerMod("migrate"))
	db := p.DB().Master()

	defer func() {
		_ = db.Close()
	}()

	err := migrator.SqliteUpFromPath(ctx, db, ".", log)
	utils.Must(err)
}

func (p *Provider) MigrationDown() {
	ctx := context.WithValue(context.Background(), migrator.ConfigKey, p.Config())
	log := migrator.NewGooseLogger(p.LoggerMod("migrate"))
	db := p.DB().Master()

	defer func() {
		_ = db.Close()
	}()

	err := migrator.SqliteResetFromPath(ctx, db, ".", log)
	utils.Must(err)
}

func (p *Provider) Scheduler() *scheduler.Scheduler {
	if p.scheduler == nil {
		var err error

		p.scheduler, err = scheduler.New(p.Config().Scheduler.StopTimeout)
		utils.MustMsg(err, "failed create scheduler")

		p.Closer().Add(func(_ context.Context) error {
			return p.scheduler.Stop()
		})
	}
	return p.scheduler
}

func (p *Provider) Telegram() *bot.Bot {
	if p.telegram == nil {
		var err error

		p.telegram, err = bot.New(p.Config().Telegram.Token)
		utils.MustMsg(err, "fail init telegram service")

		p.Closer().Add(func(ctx context.Context) error {
			_, err = p.telegram.Close(ctx)
			return err
		})
	}
	return p.telegram
}
