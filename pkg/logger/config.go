package logger

import (
	"io"
	"log/slog"
	"os"
)

type config struct {
	format  string
	options *Options
	out     io.Writer
}

func newConfig(opts ...Option) *config {
	cfg := &config{
		format:  FormatJson,
		options: &Options{Level: slog.LevelError},
		out:     os.Stdout,
	}

	for _, opt := range opts {
		opt(cfg)
	}

	return cfg
}

func (c *config) Handler() slog.Handler {
	switch c.format {
	case FormatJson:
		return slog.NewJSONHandler(c.out, c.options.ToSlogHandleOptions())
	case FormatText:
		return slog.NewTextHandler(c.out, c.options.ToSlogHandleOptions())
	case FormatPretty:
		return NewPrettyHandler(c.out, c.options)
	case FormatDiscard:
		return slog.DiscardHandler
	default:
		return slog.NewJSONHandler(c.out, c.options.ToSlogHandleOptions())
	}
}
