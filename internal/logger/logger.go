package logger

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger = logrus.New()

type CustomFormatter struct {
	logrus.JSONFormatter
}

func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	if entry.Caller != nil {
		entry.Data["func"] = entry.Caller.Function
		delete(entry.Data, "file")
	}
	return f.JSONFormatter.Format(entry)
}

func InitLogger() {
	Logger.SetFormatter(&CustomFormatter{})
	Logger.SetReportCaller(true)

	logLevel := os.Getenv("LOG_LEVEL")
	switch logLevel {
	case "debug":
		Logger.SetLevel(logrus.DebugLevel)
	case "warn":
		Logger.SetLevel(logrus.WarnLevel)
	case "error":
		Logger.SetLevel(logrus.ErrorLevel)
	default:
		Logger.SetLevel(logrus.InfoLevel)
	}

	logFile := &lumberjack.Logger{
		Filename:   "./logs/app.log",
		MaxSize:    10,
		MaxBackups: 3,
		MaxAge:     28,
		Compress:   true,
	}

	Logger.SetOutput(io.MultiWriter(os.Stdout, logFile))
}
