package logger

import (
	"io"
	"log/slog"
)

const (
	// Log levels
	LogLevelInfo  = "info"
	LogLevelDebug = "debug"
)

var defaultLogger *slog.Logger

// InitDefaultLogger initializes the default logger.
func InitDefaultLogger(w io.Writer, level string) *slog.Logger {
	defaultLogger = InitLogger(w, level)
	return DefaultLogger()
}

// InitLogger initializes a logger with the given level.
func InitLogger(w io.Writer, level string) *slog.Logger {
	if w == nil {
		w = io.Discard
	}
	lvl := new(slog.LevelVar)
	if level == LogLevelDebug {
		lvl.Set(slog.LevelDebug)
	} else if level == LogLevelInfo {
		lvl.Set(slog.LevelInfo)
	}
	return slog.New(slog.NewJSONHandler(w, &slog.HandlerOptions{
		Level: lvl,
	}))
}

func DefaultLogger() *slog.Logger {
	if defaultLogger == nil {
		InitDefaultLogger(nil, LogLevelInfo)
	}
	return defaultLogger
}
