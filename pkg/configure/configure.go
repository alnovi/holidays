package configure

import (
	"context"
	"fmt"
	"time"

	"github.com/sethvargo/go-envconfig"
)

func LoadFromEnv(config any) error {
	ctx, ctxCancel := context.WithTimeout(context.Background(), time.Second)
	defer ctxCancel()

	err := envconfig.Process(ctx, config)
	if err != nil {
		return fmt.Errorf("failed to envconfig.Process: %w", err)
	}

	return nil
}
