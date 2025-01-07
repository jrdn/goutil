package log

import "log/slog"

var (
	Any      = slog.Any
	Bool     = slog.Bool
	Duration = slog.Duration
	Float64  = slog.Float64
	Int      = slog.Int
	Int64    = slog.Int64
	String   = slog.String
	Time     = slog.Time
	Uint64   = slog.Uint64

	Group = slog.Group

	Debug = slog.Debug
	Info  = slog.Info
	Warn  = slog.Warn
	Error = slog.Error
)

const (
	LevelDebug = slog.LevelDebug
	LevelInfo  = slog.LevelInfo
	LevelWarn  = slog.LevelWarn
	LevelError = slog.LevelError
)
