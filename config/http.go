package config

type Http struct {
	Host string `env:"HOST,default=0.0.0.0"`
	Port string `env:"PORT,default=8080"`
}
