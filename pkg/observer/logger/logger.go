package logger

type Logger interface {
	Info(msg string, a ...interface{})
	Fatal(msg string, a ...interface{})
	Warn(msg string, a ...interface{})
	Error(msg string, a ...interface{})
	Panic(msg string, a ...interface{})
}
