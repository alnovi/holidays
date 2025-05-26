package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/alnovi/holidays/config"
	"github.com/alnovi/holidays/docs"
	"github.com/alnovi/holidays/internal/provider"
	"github.com/alnovi/holidays/internal/transport/http/middleware"
	"github.com/alnovi/holidays/pkg/server"
)

type App struct {
	Provider    *provider.Provider
	Controllers []server.HttpController
	HttpServer  *server.HttpServer
}

func NewApp(cfg *config.Config) *App {
	app := &App{Provider: provider.New(cfg)}

	defer func() {
		if err := recover(); err != nil {
			app.Provider.LoggerMod("app-server").Error(fmt.Sprintf("failed init app: %s", err.(error).Error()))
			os.Exit(1)
		}
	}()

	app.Provider.MigrationUp()
	app.initControllers()
	app.initHTTPServer()
	app.initSwag()

	return app
}

func (app *App) Start(ctx context.Context) {
	if ctx == nil {
		ctx = context.Background()
	}

	ctx, cancel := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)

	defer func() {
		if err := app.Provider.Closer().Close(); err != nil {
			app.Provider.LoggerMod("closer").Error(err.Error())
		}

		if err := recover(); err != nil {
			app.Provider.LoggerMod("app-server").Error(err.(error).Error())
			os.Exit(1)
		}

		cancel()
	}()

	go func() {
		app.Provider.Scheduler().Start()
	}()

	go func() {
		err := app.HttpServer.Start(app.Provider.Config().Http.Host, app.Provider.Config().Http.Port)
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			app.Provider.LoggerMod("http-server").Error(err.(error).Error())
			cancel()
		}
	}()

	app.Provider.LoggerMod("http-server").Info("server started", "port", app.Provider.Config().Http.Port)

	<-ctx.Done()
}

func (app *App) initControllers() {
	app.Controllers = []server.HttpController{}
}

func (app *App) initHTTPServer() {
	app.HttpServer = server.NewHttpServer(
		server.WithHideBanner(),
		server.WithHidePort(),
		//server.WithRender(server.NewHttpRenderFromFS(web.StaticFS, "out/html")),
		//server.WithErrorHandler(controller.NewErrorController().Handle),
		server.WithValidator(app.Provider.Validator()),
		server.WithControllers(app.Controllers...),
		server.WithCors(app.Provider.Config().IsDevelopment()),
	)

	app.HttpServer.Pre(middleware.TrailingSlash())
	app.HttpServer.Use(middleware.RequestLogger(app.Provider.LoggerMod("http-request")))

	//app.HttpServer.FileFS("/favicon.png/", "public/icon.png", web.StaticFS)
	//app.HttpServer.StaticFS("/assets/*", echo.MustSubFS(web.StaticFS, "out/assets"))
	//app.HttpServer.StaticFS("/public/*", echo.MustSubFS(web.StaticFS, "public"))
	app.HttpServer.GET("/swagger/*", echoSwagger.WrapHandler)

	app.Provider.Closer().Add(app.HttpServer.Shutdown)
}

func (app *App) initSwag() {
	docs.SwaggerInfo.Version = app.Provider.Config().App.Version
	docs.SwaggerInfo.Host = app.Provider.Config().App.Host
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}
