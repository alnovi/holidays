package config

type Database struct {
	Database string `env:"DATABASE,default=./holidays.db"`
}
