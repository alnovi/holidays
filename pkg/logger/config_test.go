package logger

import (
	"bytes"
	"io"
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigLevel(t *testing.T) {
	testCases := []struct {
		name     string
		actLevel string
		expLevel slog.Level
	}{
		{
			name:     "Debug",
			actLevel: LevelDebug,
			expLevel: slog.LevelDebug,
		}, {
			name:     "Info",
			actLevel: LevelInfo,
			expLevel: slog.LevelInfo,
		}, {
			name:     "Warn",
			actLevel: LevelWarn,
			expLevel: slog.LevelWarn,
		}, {
			name:     "Error",
			actLevel: LevelError,
			expLevel: slog.LevelError,
		}, {
			name:     "Default",
			actLevel: "",
			expLevel: slog.LevelError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cfg := newConfig(WithLevel(tc.actLevel))
			assert.Equal(t, cfg.options.Level, tc.expLevel)
		})
	}
}

func TestConfigFormat(t *testing.T) {
	testCases := []struct {
		name    string
		format  string
		handler slog.Handler
	}{
		{
			name:    "Text",
			format:  FormatText,
			handler: slog.NewTextHandler(os.Stdout, nil),
		}, {
			name:    "Json",
			format:  FormatJson,
			handler: slog.NewJSONHandler(os.Stdout, nil),
		}, {
			name:    "Discard",
			format:  FormatDiscard,
			handler: slog.DiscardHandler,
		}, {
			name:    "Pretty",
			format:  FormatPretty,
			handler: NewPrettyHandler(os.Stdout, nil),
		}, {
			name:    "Default",
			format:  "",
			handler: slog.NewJSONHandler(os.Stdout, nil),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cfg := newConfig(WithFormat(tc.format))
			assert.IsType(t, tc.handler, cfg.Handler())
		})
	}
}

func TestConfigWriter(t *testing.T) {
	testCases := []struct {
		name   string
		actOut io.Writer
		expOut io.Writer
	}{
		{
			name:   "Stdout",
			actOut: os.Stdout,
			expOut: os.Stdout,
		}, {
			name:   "Buffer",
			actOut: bytes.NewBuffer(nil),
			expOut: bytes.NewBuffer(nil),
		}, {
			name:   "Default",
			actOut: nil,
			expOut: os.Stdout,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cfg := newConfig(WithWriter(tc.actOut))
			assert.Equal(t, cfg.out, tc.expOut)
		})
	}
}

func TestConfigOptions(t *testing.T) {
	testCases := []struct {
		name   string
		actOpt *Options
		expOpt *Options
	}{
		{
			name:   "Default",
			actOpt: nil,
			expOpt: &Options{Level: slog.LevelError},
		}, {
			name:   "Custom",
			actOpt: &Options{Level: slog.LevelDebug},
			expOpt: &Options{Level: slog.LevelDebug},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cfg := newConfig(WithOptions(tc.actOpt))
			assert.Equal(t, tc.expOpt, cfg.options)
		})
	}
}
