package log

import (
	"io"
	"os"

	"github.com/rs/zerolog"
)

type Option func(lc *logConfig)

type logConfig struct {
	level   zerolog.Level
	stream  io.Writer
	console bool
	appName string
}

func WithLevel(level zerolog.Level) Option {
	return func(lc *logConfig) {
		lc.level = level
	}
}

func WithStream(stream io.Writer) Option {
	return func(lc *logConfig) {
		lc.stream = stream
	}
}

func WithStdout() Option {
	return func(lc *logConfig) {
		lc.stream = os.Stdout
	}
}

func WithStderr() Option {
	return func(lc *logConfig) {
		lc.stream = os.Stderr
	}
}

func WithConsoleOutput() Option {
	return func(lc *logConfig) {
		lc.console = true
	}
}

func WithAppName(name string) Option {
	return func(lc *logConfig) {
		lc.appName = name
	}
}
