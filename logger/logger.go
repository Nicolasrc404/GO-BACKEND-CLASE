package logger

import (
	"log"
)

type Logger struct {
	*log.Logger
}

func NewLogger() *Logger {
	return &Logger{log.New(log.Writer(), "", 0)}
}

func (l *Logger) Fatal(err error) {
	l.Fatalf("ATAL | ERROR: %v", err)
}
