package io

import (
	"io"

	"gopkg.in/natefinch/lumberjack.v2"
)

type rotateLogWriterConfig struct {
	fileName   string
	maxSize    int
	maxAge     int
	maxBackups int
	localTime  bool
	compress   bool
}

// NewRotateLogWriter is ...
func NewRotateLogWriter(c *rotateLogWriterConfig) io.Writer {
	return &lumberjack.Logger{
		Filename:   c.fileName,
		MaxSize:    c.maxSize,    // megabytes
		MaxBackups: c.maxBackups, // max backup size
		MaxAge:     c.maxAge,     //days
		LocalTime:  c.localTime,  // localtime
		Compress:   c.compress,   // disabled by default
	}
}
