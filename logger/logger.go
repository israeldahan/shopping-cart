package logger

import (
    "go.uber.org/zap"
)

var logger *zap.Logger

// Init initializes the logger
func Init(appName string) {
	logger, _ = zap.NewProduction()

	defer logger.Sync()

	logger.Info("logger initialized", zap.String("app", appName))

}

// Info logs an info message
func Info(msg string, tags ...zap.Field) {
	logger.Info(msg, tags...)
	logger.Sync()
}

// Error logs an error message
func Error(msg string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	logger.Error(msg, tags...)
	logger.Sync()
}

// Fatal logs a fatal message
func Fatal(msg string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("fatal", err))
	logger.Fatal(msg, tags...)
	logger.Sync()
}

// Debug logs a debug message
func Debug(msg string, tags ...zap.Field) {
	logger.Debug(msg, tags...)
	logger.Sync()
}

// Warn logs a warning message
func Warn(msg string, tags ...zap.Field) {
	logger.Warn(msg, tags...)
	logger.Sync()
}
