package logger

var (
	rootLogger = newLogger(NewConfig("default"))
)

// Logger is ...
type Logger interface {
	Debug(string)
	Debugf(string, ...interface{})
	Info(string)
	Infof(string, ...interface{})
	Warn(string)
	Warnf(string, ...interface{})
	Error(string)
	Errorf(string, ...interface{})
	Panic(string)
	Panicf(string, ...interface{})
}

// Debug is ...
func Debug(msg string) {
	rootLogger.Debug(msg)
}

// Debugf is ...
func Debugf(format string, args ...interface{}) {
	rootLogger.Debugf(format, args...)
}

// Info is ...
func Info(msg string) {
	rootLogger.Info(msg)
}

// Infof is ...
func Infof(format string, args ...interface{}) {
	rootLogger.Infof(format, args...)
}

// Warn is ...
func Warn(msg string) {
	rootLogger.Warn(msg)
}

// Warnf is ...
func Warnf(format string, args ...interface{}) {
	rootLogger.Warnf(format, args...)
}

// Error is ...
func Error(msg string) {
	rootLogger.Error(msg)
}

// Errorf is ...
func Errorf(format string, args ...interface{}) {
	rootLogger.Errorf(format, args...)
}

// Panic is ...
func Panic(msg string) {
	rootLogger.Panic(msg)
}

// Panicf is ...
func Panicf(format string, args ...interface{}) {
	rootLogger.Panicf(format, args...)
}

// SetupRootLogger is ...
func SetupRootLogger(configs ...*Config) {
	rootLogger = newLogger(configs...)
}

func newLogger(configs ...*Config) Logger {
	switch len(configs) {
	case 0:
		panic("configration not found")
	case 1:
		return newStandardLogger(configs[0])
	default:
		return newMultiLogger(configs...)
	}
}
