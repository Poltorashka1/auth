package logger

import (
	"auth/internal/logger/lib"
	"fmt"
	"log/slog"
	"os"
)

type SlogLogger struct {
	log *slog.Logger
}

func Load() Logger {
	opts := colorLog.PrettyHandlerOptions{
		SlogOpts: slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := colorLog.NewPrettyHandler(os.Stdout, opts)

	return &SlogLogger{log: slog.New(handler)}
}

func (l *SlogLogger) Info(err string) {
	l.log.Info(err)
}

func (l *SlogLogger) Error(err error) {
	l.log.Error(err.Error())
}

func (l *SlogLogger) ErrorOp(op string, err error) {
	l.log.Error(fmt.Sprintf("%s: %s", op, err.Error()))
}

func (l *SlogLogger) Fatal(err error) {
	l.log.Error(err.Error())
	os.Exit(1)
}

func (l *SlogLogger) FatalOp(op string, err error) {
	l.log.Error(fmt.Sprintf("%s: %s", op, err.Error()))
	os.Exit(1)
}
