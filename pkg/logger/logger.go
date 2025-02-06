package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
	"smsc/internal/config"
)

// Config holds logger configuration
type Config struct {
	Level      string
	Format     string
	Output     string
	FilePath   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}

// New creates a new configured logger
func New(cfg config.LoggingConfig) (*logrus.Logger, error) {
	log := logrus.New()

	// Set log level
	level, err := logrus.ParseLevel(strings.ToLower(cfg.Level))
	if err != nil {
		return nil, fmt.Errorf("invalid log level: %w", err)
	}
	log.SetLevel(level)

	// Set log format
	switch strings.ToLower(cfg.Format) {
	case "json":
		log.SetFormatter(&logrus.JSONFormatter{})
	case "text":
		log.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
	default:
		return nil, fmt.Errorf("unsupported log format: %s", cfg.Format)
	}

	// Set log output
	switch strings.ToLower(cfg.Output) {
	case "stdout":
		log.SetOutput(os.Stdout)
	case "stderr":
		log.SetOutput(os.Stderr)
	case "file":
		if err := os.MkdirAll(filepath.Dir(cfg.FilePath), 0755); err != nil {
			return nil, err
		}
		file, err := os.OpenFile(cfg.FilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return nil, fmt.Errorf("failed to open log file: %w", err)
		}
		log.SetOutput(file)
	default:
		return nil, fmt.Errorf("unsupported log output: %s", cfg.Output)
	}

	return log, nil
}

// Fields type, used to pass to `WithFields`.
type Fields logrus.Fields

// WithFields creates an entry from the standard logger and adds multiple fields to it
func WithFields(fields Fields) *logrus.Entry {
	return logrus.WithFields(logrus.Fields(fields))
}

// Debug logs a message at level Debug on the standard logger
func Debug(args ...interface{}) {
	logrus.Debug(args...)
}

// Info logs a message at level Info on the standard logger
func Info(args ...interface{}) {
	logrus.Info(args...)
}

// Warn logs a message at level Warn on the standard logger
func Warn(args ...interface{}) {
	logrus.Warn(args...)
}

// Error logs a message at level Error on the standard logger
func Error(args ...interface{}) {
	logrus.Error(args...)
}

// Fatal logs a message at level Fatal on the standard logger then the process will exit with status set to 1
func Fatal(args ...interface{}) {
	logrus.Fatal(args...)
}

// Panic logs a message at level Panic on the standard logger
func Panic(args ...interface{}) {
	logrus.Panic(args...)
} 