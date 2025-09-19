package logger

import (
	"log/slog"
	"os"
)

// Init initializes the logger ands sets it as the default logger globally.
func Init() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	// set as the default logger globally
	slog.SetDefault(logger)

	logger.Info("logger initialized")
}

// Info logs an info level log.
func Info(msg string, args ...any) {
	slog.Info(msg, args...)
}

// Error logs an error level log.
func Error(msg string, args ...any) {
	slog.Error(msg, args...)
}

// Fatal logs a fatal error and exits the program.
func Fatal(message string, err error) {
	slog.Error(message, slog.String("error", err.Error()))
	os.Exit(1)
}
