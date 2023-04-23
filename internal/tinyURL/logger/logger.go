package logger

import (
	"github.com/sirupsen/logrus"
)

var MainLogger *logrus.Entry

func init() {
	MainLogger = logrus.NewEntry(logrus.StandardLogger())
	MainLogger.Logger.SetFormatter(&logrus.TextFormatter{ForceColors: true})
}
