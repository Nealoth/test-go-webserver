package util

import (
	"github.com/sirupsen/logrus"
	"restapi/internal/configuration"
)

var logger *logrus.Logger

func Logger() *logrus.Logger {
	if logger == nil {
		logger = initLogger()
	}

	return logger
}

func initLogger() *logrus.Logger {
	logger := logrus.New()
	configureLogger(logger)
	return logger
}

func configureLogger(logger *logrus.Logger) *logrus.Logger {
	loggingConf := configuration.ConfigurationManager().GetCached().Logging

	lvl, err := logrus.ParseLevel(loggingConf.Severity)

	if err != nil {
		panic("Incorrect severity level: " + loggingConf.Severity)
	}

	logger.SetLevel(lvl)

	tf := &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp: true,
	}

	logger.SetFormatter(tf)

	return logger
}
