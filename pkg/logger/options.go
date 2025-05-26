package logger

import "log/slog"

type ReplaceAttrFunc func(groups []string, attr slog.Attr) slog.Attr

type Options struct {
	// Level is the minimum [slog.Level] that will be logged.
	// Records with lower levels will be discarded.
	Level slog.Leveler

	// ReplaceAttr is used to rewrite each non-group [slog.Attr] before it is
	// logged. See https://pkg.go.dev/log/slog#HandlerOptions for details.
	ReplaceAttr ReplaceAttrFunc

	// AddSource enables computing the source code position of the log
	// statement and adds [slog.SourceKey] attributes to the output.
	AddSource bool

	// DisableColor disables the use of ANSI colour codes in messages.
	DisableColor bool

	// TimeFormatter is the [time.Time] formatter used to format log timestamps.
	TimeFormatter TimeFormatter

	// LevelFormatter is the [slog.Level] formatter used to format log levels.
	LevelFormatter LevelFormatter

	// SourceFormatter is the [slog.Source] formatter used to format log sources.
	SourceFormatter SourceFormatter
}

func (o *Options) ToSlogHandleOptions() *slog.HandlerOptions {
	return &slog.HandlerOptions{
		Level:       o.Level,
		ReplaceAttr: o.ReplaceAttr,
		AddSource:   o.AddSource,
	}
}
