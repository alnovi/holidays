package logger

import (
	"log/slog"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDefaultTimeFormatter(t *testing.T) {
	now := time.Now()

	testCases := []struct {
		name   string
		expect string
	}{
		{
			name:   "DateTime",
			expect: now.Format(time.DateTime),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			buf := newBuffer()
			DefaultTimeFormatter(time.DateTime)(buf, now)
			assert.Equal(t, tc.expect, buf.String())
		})
	}
}

func TestDefaultLevelFormatter(t *testing.T) {
	testCases := []struct {
		name   string
		color  bool
		level  slog.Level
		expect string
	}{
		{
			name:   "Debug",
			color:  false,
			level:  slog.LevelDebug,
			expect: "DBG",
		}, {
			name:   "Debug color",
			color:  true,
			level:  slog.LevelDebug,
			expect: ansiLevelDebug + "DBG" + ansiReset,
		}, {
			name:   "Info",
			color:  false,
			level:  slog.LevelInfo,
			expect: "INF",
		}, {
			name:   "Info color",
			color:  true,
			level:  slog.LevelInfo,
			expect: ansiLevelInfo + "INF" + ansiReset,
		}, {
			name:   "Warn",
			color:  false,
			level:  slog.LevelWarn,
			expect: "WRN",
		}, {
			name:   "Warn color",
			color:  true,
			level:  slog.LevelWarn,
			expect: ansiLevelWarn + "WRN" + ansiReset,
		}, {
			name:   "Error",
			color:  false,
			level:  slog.LevelError,
			expect: "ERR",
		}, {
			name:   "Error color",
			color:  true,
			level:  slog.LevelError,
			expect: ansiLevelError + "ERR" + ansiReset,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			buf := newBuffer()
			DefaultLevelFormatter(tc.color)(buf, tc.level)
			assert.Equal(t, tc.expect, buf.String())
		})
	}
}

func TestDefaultSourceFormatter(t *testing.T) {
	testCases := []struct {
		name   string
		color  bool
		source *slog.Source
		expect string
	}{
		{
			name:   "Source nil",
			color:  false,
			source: nil,
			expect: "",
		}, {
			name:   "Source nil color",
			color:  true,
			source: nil,
			expect: "",
		}, {
			name:   "Source main",
			color:  false,
			source: &slog.Source{Function: "main", File: "main.go", Line: 21},
			expect: "<main.go:21>",
		}, {
			name:   "Source main color",
			color:  true,
			source: &slog.Source{Function: "main", File: "main.go", Line: 21},
			expect: ansiFaint + "<main.go:21>" + ansiReset,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			buf := newBuffer()
			DefaultSourceFormatter(tc.color)(buf, tc.source)
			assert.Equal(t, tc.expect, buf.String())
		})
	}
}

func TestAppendLevelDelta(t *testing.T) {
	testCases := []struct {
		name   string
		delta  slog.Level
		expect string
	}{
		{
			name:   "Debug",
			delta:  slog.LevelDebug,
			expect: "--4",
		}, {
			name:   "Info",
			delta:  slog.LevelInfo,
			expect: "",
		}, {
			name:   "Warn",
			delta:  slog.LevelWarn,
			expect: "+4",
		}, {
			name:   "Err",
			delta:  slog.LevelError,
			expect: "+8",
		}, {
			name:   "Int +",
			delta:  slog.Level(21),
			expect: "+21",
		}, {
			name:   "Int -",
			delta:  slog.Level(-21),
			expect: "--21",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			buf := newBuffer()
			appendLevelDelta(buf, tc.delta)
			assert.Equal(t, tc.expect, buf.String())
		})
	}
}
