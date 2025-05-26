package config

import "time"

type Scheduler struct {
	StopTimeout time.Duration `env:"STOP_TIMEOUT,default=5s"`
}
