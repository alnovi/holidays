package closer

import (
	"context"
	"fmt"
	"slices"
	"strings"
	"sync"
	"time"
)

type CloseFunc func(ctx context.Context) error

type Closer struct {
	mu sync.Mutex
	t  time.Duration
	fs []CloseFunc
}

func New(t time.Duration) *Closer {
	return &Closer{t: t}
}

func (c *Closer) Add(f CloseFunc) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.fs = append(c.fs, f)
}

func (c *Closer) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	ctx, cancel := context.WithTimeout(context.Background(), c.t)
	defer cancel()

	slices.Reverse(c.fs)
	defer func() {
		c.fs = nil
	}()

	messages := make([]string, 0, len(c.fs))
	complete := make(chan struct{}, 1)

	go func() {
		for _, f := range c.fs {
			if err := f(ctx); err != nil {
				messages = append(messages, err.Error())
			}
		}
		complete <- struct{}{}
	}()

	select {
	case <-complete:
		break
	case <-ctx.Done():
		return fmt.Errorf("shutdown cancelled: %s", ctx.Err())
	}

	if len(messages) > 0 {
		return fmt.Errorf("shutdown finished with error(s): %s", strings.Join(messages, "; "))
	}

	return nil
}
