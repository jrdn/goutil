package log

import (
	"io"
	"log/slog"
)

type logFormat string

var (
	FormatJSON logFormat = "json"
	FormatText logFormat = "text"
)

type logConfig struct {
	stream io.Writer
	level  *slog.LevelVar
	format logFormat
}

type Option func(*logConfig)

/* WithWriter sets the log writer. */
func WithWriter(w io.Writer) Option {
	return func(config *logConfig) {
		config.stream = w
	}
}

/* WithLevel sets the log level. */
func WithLevel(level slog.Level) Option {
	return func(config *logConfig) {
		config.level.Set(level)
	}
}

/* WithFormat sets the log format. */
func WithFormat(format logFormat) Option {
	return func(config *logConfig) {
		config.format = format
	}
}

/* SetLevel sets the log level. */
func SetLevel(level slog.Level) {
	config.level.Set(level)
}
