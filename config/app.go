package config

import "time"

type App struct {
	Environment string        `env:"ENVIRONMENT,default=production"`
	Version     string        `env:"VERSION,default=0.0.0"`
	Host        string        `env:"HOST"`
	Shutdown    time.Duration `env:"SHUTDOWN,default=10s"`
}
