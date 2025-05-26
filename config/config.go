package config

import "slices"

const (
	AppEnvironmentProduction  = "production"
	AppEnvironmentDevelopment = "development"
	AppEnvironmentTesting     = "testing"
)

type Config struct {
	App       App       `env:",prefix=APP_"`
	Logger    Logger    `env:",prefix=LOG_"`
	Http      Http      `env:",prefix=HTTP_"`
	Database  Database  `env:",prefix=DB_"`
	Telegram  Telegram  `env:",prefix=TELEGRAM_"`
	Scheduler Scheduler `env:",prefix=SCHEDULER_"`
}

func (c *Config) IsProduction() bool {
	return c.App.Environment == AppEnvironmentProduction || !slices.Contains([]string{
		AppEnvironmentDevelopment,
		AppEnvironmentTesting,
	}, c.App.Environment)
}

func (c *Config) IsDevelopment() bool {
	return c.App.Environment == AppEnvironmentDevelopment
}

func (c *Config) Normalize() {}
