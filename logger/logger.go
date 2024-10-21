package logger

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

var (
	Logger *logrus.Logger // Global logger
)

func init() {
	Logger = logrus.New()
	Logger.SetFormatter(&logrus.JSONFormatter{})
	Logger.SetOutput(os.Stdout)

	logLevel := logrus.InfoLevel
	if envLogLevel := os.Getenv("LOG_LEVEL"); envLogLevel != "" {
		switch strings.ToLower(envLogLevel) {
		case "debug":
			logLevel = logrus.DebugLevel
		case "info":
			logLevel = logrus.InfoLevel
		case "warn":
			logLevel = logrus.WarnLevel
		case "error":
			logLevel = logrus.ErrorLevel
		case "fatal":
			logLevel = logrus.FatalLevel
		case "panic":
			logLevel = logrus.PanicLevel
		default:
			logLevel = logrus.WarnLevel
		}
	}
	Logger.SetLevel(logLevel)
	Logger.Infof("Setting log level to: '%s;", logLevel)
}
