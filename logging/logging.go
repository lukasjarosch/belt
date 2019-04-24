package logging

import (
	"github.com/sirupsen/logrus"
	"os"
)

// NewLogger is a simple convenience wrapper to create a new pre-configured logrus.Logger
func NewLogger(debug bool, pretty bool) *logrus.Logger {
	logger := logrus.New()
	logger.SetOutput(os.Stderr)
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(&logrus.JSONFormatter{})

	if debug {
		logger.SetLevel(logrus.DebugLevel)
	}

	if pretty {
		logger.SetFormatter(&logrus.TextFormatter{})
	}

	return logger
}

// NewFromEnvironment is another convenience wrapper which reads the 'debug' and 'pretty' values from environment variables
// If the env variables do not exist, it falls back to debug=false, pretty=false
func NewFromEnvironment() *logrus.Logger {
	debug := getBoolEnv("LOG_DEBUG")
	pretty := getBoolEnv("LOG_PRETTY")

	return NewLogger(debug, pretty)
}

func getBoolEnv(key string) bool {
	value := os.Getenv(key)
	if len(value) == 0 {
		return false
	}
	return true
}
