package logger

import (
	"fmt"
	"strings"
)

// Level is ...
type Level uint8

func (level Level) String() string {
	switch level {
	case Levels.Debug:
		return "debug"
	case Levels.Info:
		return "info"
	case Levels.Warn:
		return "warn"
	case Levels.Error:
		return "error"
	case Levels.Panic:
		return "panic"
	case Levels.Fatal:
		return "fatal"
	default:
		return "unknown"
	}
}

// Levels is ...
var Levels = func() struct {
	Debug Level
	Info  Level
	Warn  Level
	Error Level
	Panic Level
	Fatal Level
} {
	const (
		Debug Level = iota + 1
		Info
		Warn
		Error
		Panic
		Fatal
	)
	return struct {
		Debug Level
		Info  Level
		Warn  Level
		Error Level
		Panic Level
		Fatal Level
	}{
		Debug: Debug,
		Info:  Info,
		Warn:  Warn,
		Error: Error,
		Panic: Panic,
		Fatal: Fatal,
	}
}()

// ParseLevel is ...
func ParseLevel(name string) Level {
	switch strings.ToLower(name) {
	case Levels.Debug.String():
		return Levels.Debug
	case Levels.Info.String():
		return Levels.Info
	case Levels.Warn.String():
		return Levels.Warn
	case Levels.Error.String():
		return Levels.Error
	case Levels.Panic.String():
		return Levels.Panic
	case Levels.Fatal.String():
		return Levels.Fatal
	default:
		panic(fmt.Sprintf("invalid log level [%s]", name))
	}
}
