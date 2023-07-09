package logger

type LoggerRepository interface {
	Debug(string)
	Info(string)
	Warn(string)
	Error(string)
}
