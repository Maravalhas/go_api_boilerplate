package logger

import (
	"api/internal/config"
	"fmt"
	"time"
)

type Logger struct {
	prefix string
}

func New(prefix string) *Logger {
	return &Logger{
		prefix: prefix,
	}
}

func (l *Logger) Info(message any) {
	l.log("INFO", message)
}

func (l *Logger) Infof(format string, args ...any) {
	l.log("INFO", fmt.Sprintf(format, args...))
}

func (l *Logger) Error(message any) {
	l.log("ERROR", message)
}

func (l *Logger) Errorf(format string, args ...any) {
	l.log("ERROR", fmt.Sprintf(format, args...))
}

func (l *Logger) Debug(message any) {
	if config.IsDebug() {
		l.log("DEBUG", message)
	}
}

func (l *Logger) Debugf(format string, args ...any) {
	if config.IsDebug() {
		l.log("DEBUG", fmt.Sprintf(format, args...))
	}
}

func (l *Logger) Fatal(message any) {
	l.log("FATAL", message)
	panic(message)
}

func (l *Logger) Fatalf(format string, args ...any) {
	l.log("FATAL", fmt.Sprintf(format, args...))
	panic(fmt.Sprintf(format, args...))
}

func (l *Logger) log(logType string, message any) {
	fmt.Printf("%v [%v] (%v): %v \n", time.Now().Format("2006-01-02 15:04:05"), logType, l.prefix, message)
}
