package logger

import (
	"github.com/sirupsen/logrus"
)

type standardLogger struct {
	*logrus.Logger
	config *Config
}

func (l *standardLogger) Debug(msg string) {
	if l.enabledLogLevel(Levels.Debug) {
		l.Logger.Debug(msg)
	}
}

func (l *standardLogger) Debugf(format string, args ...interface{}) {
	if l.enabledLogLevel(Levels.Debug) {
		l.Logger.Debugf(format, args...)
	}
}

func (l *standardLogger) Info(msg string) {
	if l.enabledLogLevel(Levels.Info) {
		l.Logger.Info(msg)
	}
}

func (l *standardLogger) Infof(format string, args ...interface{}) {
	if l.enabledLogLevel(Levels.Info) {
		l.Logger.Infof(format, args...)
	}
}

func (l *standardLogger) Warn(msg string) {
	if l.enabledLogLevel(Levels.Warn) {
		l.Logger.Warn(msg)
	}
}

func (l *standardLogger) Warnf(format string, args ...interface{}) {
	if l.enabledLogLevel(Levels.Warn) {
		l.Logger.Warnf(format, args...)
	}
}

func (l *standardLogger) Error(msg string) {
	if l.enabledLogLevel(Levels.Error) {
		l.Logger.Error(msg)
	}
}

func (l *standardLogger) Errorf(format string, args ...interface{}) {
	if l.enabledLogLevel(Levels.Error) {
		l.Logger.Errorf(format, args...)
	}
}

func (l *standardLogger) Panic(msg string) {
	if l.enabledLogLevel(Levels.Panic) {
		l.Logger.Panic(msg)
	}
}

func (l *standardLogger) Panicf(format string, args ...interface{}) {
	if l.enabledLogLevel(Levels.Panic) {
		l.Logger.Panicf(format, args...)
	}
}

func (l *standardLogger) enabledLogLevel(level Level) bool {
	return l.config.level <= level && level <= l.config.maxLevel
}

func newStandardLogger(config *Config) Logger {
	var l = logrus.New()
	{
		l.Level, _ = logrus.ParseLevel(config.level.String())
		l.Out = config.out
	}
	return &standardLogger{
		Logger: l,
		config: config,
	}
}
