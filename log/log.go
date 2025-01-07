package log

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var config = &logConfig{
	level:  &slog.LevelVar{},
	format: FormatJSON,
	stream: os.Stderr,
}

// Setup configures the logger with the given options.
func Setup(options ...Option) error {
	for _, opt := range options {
		opt(config)
	}

	var handler slog.Handler

	opts := &slog.HandlerOptions{
		Level:     config.level,
		AddSource: true,
	}

	switch config.format {
	case FormatJSON:
		handler = slog.NewJSONHandler(config.stream, opts)
	case FormatText:
		handler = slog.NewTextHandler(config.stream, opts)
	default:
		return fmt.Errorf("unknown log format: %s", config.format)
	}
	logger := slog.New(handler)
	slog.SetDefault(logger)
	return nil
}

// Get returns a logger with the given name.
func Get(name string) *slog.Logger {
	return slog.Default().With(slog.String("name", name))
}

// GetCmdLogger returns a logger with the name of the given cobra command
func GetCmdLogger(cmd *cobra.Command) *slog.Logger {
	return Get(getCmdLoggerName(cmd))
}

func getCmdLoggerName(cmd *cobra.Command) string {
	var names []string
	cmd.VisitParents(func(c *cobra.Command) {
		names = append(names, c.Name())
	})
	return strings.Join(names, ".")
}
