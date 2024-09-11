package logger

import (
	"context"
	"fmt"
	"log"
)

type Logger interface {
	Debug(args ...any)
	Info(args ...any)
	Error(args ...any)
	WithFields(key, value any) Logger
}

type logger struct {
	ctx context.Context
}

func New(ctx context.Context) Logger {
	return logger{ctx}
}

func (l logger) WithFields(key, value any) Logger {
	prefix := log.Prefix()
	log.SetPrefix(fmt.Sprintf("%v %v: %v ", prefix, key, value))
	return l
}

func (l logger) Debug(args ...any) {
	prefix := log.Prefix()
	prefix = "DEBUG: " + prefix
	logText(prefix, args)

}

func (l logger) Info(args ...any) {
	prefix := log.Prefix()
	prefix = "INFO: " + prefix
	logText(prefix, args)

}

func (l logger) Error(args ...any) {
	prefix := log.Prefix()
	prefix = "ERROR: " + prefix
	logText(prefix, args)
}

func logText(prefix string, args any) {
	log.SetPrefix(prefix + "- ")
	log.Println(args)
}
