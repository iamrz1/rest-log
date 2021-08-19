package rest_log

type LogLevel string

const (
	Info  LogLevel = "Info"
	Warn  LogLevel = "Warn"
	Debug LogLevel = "Debug"
	Error LogLevel = "Error"
	Fatal LogLevel = "Fatal"
)

type Logger interface {
	Info(fn, tid string, msg string)
	InfoPretty(fn, tid string, msg string)
	Warn(fn, tid string, msg string)
	WarnPretty(fn, tid string, msg string)
	Error(fn, tid string, msg string)
	ErrorPretty(fn, tid string, msg string)
	Print(level LogLevel, fn, tid string, msg string)
}

func New(service string, verbose bool) Logger {
	logger := NewZeroLevelLogger(service,verbose)
	logger.Info("GetDefaultStructLogger", "init", "Running in verbose mode")
	return logger
}
