package migrator

import (
	"fmt"
	"log/slog"
	"strings"
)

type GooseLogger struct {
	logger *slog.Logger
}

func NewGooseLogger(logger *slog.Logger) *GooseLogger {
	return &GooseLogger{logger: logger}
}

func (l *GooseLogger) Fatalf(format string, v ...interface{}) {
	msg := l.clearMsg(fmt.Sprintf(format, v...))
	l.logger.Error(msg)
}

func (l *GooseLogger) Printf(format string, v ...interface{}) {
	msg := l.clearMsg(fmt.Sprintf(format, v...))
	l.logger.Info(msg)
}

func (l *GooseLogger) clearMsg(msg string) string {
	msg = strings.TrimPrefix(msg, "goose: ")
	msg = strings.TrimSuffix(msg, "\n")
	return msg
}
