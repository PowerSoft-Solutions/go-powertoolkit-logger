package ports

type PowerLogger interface {
	Debug(v ...any)
	Info(v ...any)
	Warn(v ...any)
	Error(v ...any)
	Panic(v ...any)
	SetLoggerFile(logName, logDir string) error
}
