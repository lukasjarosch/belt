package logging

import (
	"github.com/sirupsen/logrus"
	"os"
)

type Logger struct {
}

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
