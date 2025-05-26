package logger

import (
	"io"
	"log/slog"
)

const (
	FormatJson    = "json"
	FormatText    = "text"
	FormatPretty  = "pretty"
	FormatDiscard = "discard"

	LevelDebug = "debug"
	LevelInfo  = "info"
	LevelWarn  = "warn"
	LevelError = "error"
)

type Option func(c *config)

func New(opts ...Option) *slog.Logger {
	return slog.New(newConfig(opts...).Handler())
}

func WithLevel(level string) Option {
	return func(c *config) {
		switch level {
		case LevelDebug:
			c.options.Level = slog.LevelDebug
		case LevelInfo:
			c.options.Level = slog.LevelInfo
		case LevelWarn:
			c.options.Level = slog.LevelWarn
		case LevelError:
			c.options.Level = slog.LevelError
		default:
			c.options.Level = slog.LevelError
		}
	}
}

func WithFormat(format string) Option {
	return func(c *config) {
		c.format = format
	}
}

func WithWriter(w io.Writer) Option {
	return func(c *config) {
		if w != nil {
			c.out = w
		}
	}
}

func WithOptions(opts *Options) Option {
	return func(c *config) {
		if opts != nil {
			c.options = opts
		}
	}
}
