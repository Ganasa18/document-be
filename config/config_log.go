package appconfig

import (
	"fmt"
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

func isProd() bool {
	return os.Getenv("APP_ENV") == "production"
}

func InitLog() {
	lv := os.Getenv("LOG_LEVEL_DEV")
	level := logrus.InfoLevel
	switch lv {
	case "PanicLevel":
		level = logrus.PanicLevel
	case "FatalLevel":
		level = logrus.FatalLevel
	case "ErrorLevel":
		level = logrus.ErrorLevel
	case "WarnLevel":
		level = logrus.WarnLevel
	case "InfoLevel":
		level = logrus.InfoLevel
	case "DebugLevel":
		level = logrus.DebugLevel
	case "TraceLevel":
		level = logrus.TraceLevel
	default:
	}

	if isProd() {
		logrus.SetFormatter(&logrus.JSONFormatter{})
		logrus.SetLevel(logrus.WarnLevel)
		logrus.SetOutput(os.Stdout)
	} else {
		logrus.SetFormatter(&logrus.JSONFormatter{})

		if lv == "" && os.Getenv("APP_DEBUG") == "True" {
			level = logrus.DebugLevel
		}
		logrus.SetLevel(level)

		if os.Getenv("DEV_FILE_LOG") == "True" {
			logfile, err := os.OpenFile("log/response.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
			if err != nil {
				fmt.Printf("error opening file : %v", err)
			}

			mw := io.MultiWriter(os.Stdout, logfile)
			logrus.SetOutput(mw)
		} else {
			logrus.SetOutput(os.Stdout)
		}
	}
}
