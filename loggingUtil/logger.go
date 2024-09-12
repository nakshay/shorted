package loggingUtil

import (
	"context"
	"fmt"
	"log"
)

type Logger interface {
	Debug(args ...any)
	Info(args ...any)
	Error(args ...any)
	Debugf(format string, args ...any)
	Infof(format string, args ...any)
	Errorf(format string, args ...any)
	WithFields(key, value any) Logger
}

type logger struct {
	ctx    context.Context
	prefix string
}

func GetLogger(ctx context.Context) Logger {
	return logger{ctx: ctx}
}

func (l logger) WithFields(key, value any) Logger {
	prefix := l.prefix
	l.prefix = fmt.Sprintf("%v %v: %v ", prefix, key, value)
	return l
}

func (l logger) Debug(args ...any) {
	prefix := fmt.Sprintf("DEBUG : %v - ", l.prefix)
	logText(prefix, args)

}

func (l logger) Info(args ...any) {
	prefix := fmt.Sprintf("INFO : %v - ", l.prefix)
	logText(prefix, args)

}

func (l logger) Error(args ...any) {
	prefix := fmt.Sprintf("ERROR : %v - ", l.prefix)
	logText(prefix, args)
}

func (l logger) Debugf(format string, args ...any) {
	prefix := fmt.Sprintf("DEBUG : %v - ", l.prefix)
	logText(prefix, fmt.Sprintf(format, args))

}

func (l logger) Infof(format string, args ...any) {
	prefix := fmt.Sprintf("INFO : %v - ", l.prefix)
	logText(prefix, fmt.Sprintf(format, args))

}

func (l logger) Errorf(format string, args ...any) {
	prefix := fmt.Sprintf("ERROR : %v - ", l.prefix)
	logText(prefix, fmt.Sprintf(format, args))
}

func logText(prefix string, args any) {
	log.SetPrefix(prefix)
	log.Println(args)
}
