package logger

type multiLogger struct {
	loggers []Logger
}

func (l *multiLogger) Debug(msg string) {
	for _, logger := range l.loggers {
		logger.Debug(msg)
	}
}

func (l *multiLogger) Debugf(format string, args ...interface{}) {
	for _, logger := range l.loggers {
		logger.Debugf(format, args...)
	}
}

func (l *multiLogger) Info(msg string) {
	for _, logger := range l.loggers {
		logger.Info(msg)
	}
}

func (l *multiLogger) Infof(format string, args ...interface{}) {
	for _, logger := range l.loggers {
		logger.Infof(format, args...)
	}
}

func (l *multiLogger) Warn(msg string) {
	for _, logger := range l.loggers {
		logger.Warn(msg)
	}
}

func (l *multiLogger) Warnf(format string, args ...interface{}) {
	for _, logger := range l.loggers {
		logger.Warnf(format, args...)
	}
}

func (l *multiLogger) Error(msg string) {
	for _, logger := range l.loggers {
		logger.Error(msg)
	}
}

func (l *multiLogger) Errorf(format string, args ...interface{}) {
	for _, logger := range l.loggers {
		logger.Errorf(format, args...)
	}
}

func (l *multiLogger) Panic(msg string) {
	for _, logger := range l.loggers {
		logger.Panic(msg)
	}
}

func (l *multiLogger) Panicf(format string, args ...interface{}) {
	for _, logger := range l.loggers {
		logger.Panicf(format, args...)
	}
}

func newMultiLogger(configs ...*Config) Logger {
	var loggers = make([]Logger, len(configs))
	for i, config := range configs {
		loggers[i] = newStandardLogger(config)
	}
	return &multiLogger{
		loggers: loggers,
	}
}
