package rest_log

import (
	"github.com/rs/zerolog"
	"os"
)

type zeroLevelLogger struct {
	StdLog  zerolog.Logger
	ErrLog  zerolog.Logger
	verbose bool
}

func NewZeroLevelLogger(verbose bool) Logger {
	stdLog := zerolog.New(os.Stdout).With().CallerWithSkipFrameCount(3).Str("service", "auth").Logger()
	errLog := zerolog.New(os.Stdout).With().CallerWithSkipFrameCount(3).Stack().Str("service", "auth").Logger()

	return &zeroLevelLogger{
		StdLog:  stdLog,
		ErrLog:  errLog,
		verbose: verbose,
	}
}
func (l zeroLevelLogger) Infoln(fn, tid string, msg string) {
	level := zerolog.Disabled
	if l.verbose {
		level = zerolog.InfoLevel
	}
	l.StdLog.WithLevel(level).Str("function", fn).Str("tid", tid).Msg(msg)
}

func (l zeroLevelLogger) Infof(fn, tid string, format string, args ...interface{}) {
	level := zerolog.Disabled
	if l.verbose {
		level = zerolog.InfoLevel
	}
	l.StdLog.WithLevel(level).Str("function", fn).Str("tid", tid).Msgf(format, args...)
}

func (l zeroLevelLogger) Warnln(fn, tid string, msg string) {
	level := zerolog.Disabled
	if l.verbose {
		level = zerolog.WarnLevel
	}
	l.StdLog.WithLevel(level).Str("function", fn).Str("tid", tid).Msg(msg)
}

func (l zeroLevelLogger) Errorln(fn, tid string, msg string) {
	level := zerolog.Disabled
	if l.verbose {
		level = zerolog.ErrorLevel
	}
	l.ErrLog.WithLevel(level).Str("function", fn).Str("tid", tid).Msg(msg)
}
func (l zeroLevelLogger) Errorf(fn, tid string, format string, args ...interface{}) {
	level := zerolog.Disabled
	if l.verbose {
		level = zerolog.ErrorLevel
	}
	l.ErrLog.WithLevel(level).Str("function", fn).Str("tid", tid).Msgf(format, args...)
}

func (l zeroLevelLogger) Print(level LogLevel, fn, tid string, msg string) {
	l.StdLog.Log().Str("level", getZeroLevel(level).String()).Str("function", fn).Str("tid", tid).Msg(msg)
}

func getZeroLevel(level LogLevel) zerolog.Level {
	switch level {
	case Info:
		return zerolog.InfoLevel
	case Warn:
		return zerolog.WarnLevel
	case Debug:
		return zerolog.DebugLevel
	case Error:
		return zerolog.ErrorLevel
	case Fatal:
		return zerolog.FatalLevel
	default:
		return zerolog.InfoLevel
	}
}
