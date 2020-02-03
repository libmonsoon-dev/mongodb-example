package usecases

type Logger interface {
	Info(v ...interface{})
	Infof(format string, v ...interface{})
	Infoln(v ...interface{})
	Warning(v ...interface{})
	Warningf(format string, v ...interface{})
	Warningln(v ...interface{})
	Error(v ...interface{})
	Errorf(format string, v ...interface{})
	Errorln(v ...interface{})
}
