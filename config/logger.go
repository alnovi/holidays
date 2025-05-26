package config

type Logger struct {
	Format string `env:"FORMAT,default=json"`
	Level  string `env:"LEVEL,default=error"`
}
