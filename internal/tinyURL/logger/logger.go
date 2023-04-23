package logger

import (
	"github.com/sirupsen/logrus"
)

var MainLogger *Logger

type LoggerInterface interface {
	LogInfo(data interface{})
	LogError(data interface{})
}

type Logger struct {
	Logger *logrus.Entry
}

func (l *Logger) LogInfo(data interface{}) {
	l.Logger.Info(data)
}

func (l *Logger) LogError(data interface{}) {
	l.Logger.Error(data)
}
