package infrastructure

import (
	"mongodb-example/src/usecases"
	"os"

	"github.com/google/logger"
)

type loggerWrapper struct {
	log *logger.Logger
	err *logger.Logger
}

func (l *loggerWrapper) Info(v ...interface{}) {
	l.log.Info(v...)
}
func (l *loggerWrapper) Infof(format string, v ...interface{}) {
	l.log.Infof(format, v...)
}
func (l *loggerWrapper) Infoln(v ...interface{}) {
	l.log.Infoln(v...)
}
func (l *loggerWrapper) Warning(v ...interface{}) {
	l.err.Warning(v...)
}
func (l *loggerWrapper) Warningf(format string, v ...interface{}) {
	l.err.Warningf(format, v...)
}
func (l *loggerWrapper) Warningln(v ...interface{}) {
	l.err.Warningln(v...)
}
func (l *loggerWrapper) Error(v ...interface{}) {
	l.err.Error(v...)
}
func (l *loggerWrapper) Errorf(format string, v ...interface{}) {
	l.err.Errorf(format, v...)
}
func (l *loggerWrapper) Errorln(v ...interface{}) {
	l.err.Errorln(v...)
}

func NewStdLogger() usecases.Logger {
	return &loggerWrapper{
		logger.Init("INFO", true, false, os.Stdout),
		logger.Init("ERROR", false, false, os.Stderr),
	}
}
