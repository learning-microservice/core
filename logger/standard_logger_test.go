package logger

import (
	"bytes"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestStandardLogger(t *testing.T) {
	tests := []struct {
		name     string
		config   *Config
		logFunc  func(logger Logger)
		expected string
	}{
		// write log test patterns
		{
			config: NewConfig("test"),
			logFunc: func(logger Logger) {
				logger.Debug("debug message")
			},
			expected: `level=debug msg="debug message"`,
		},
		{
			config: NewConfig("test"),
			logFunc: func(logger Logger) {
				logger.Debugf("debug message [%d]", 111)
			},
			expected: `level=debug msg="debug message [111]"`,
		},
		{
			config: NewConfig("test"),
			logFunc: func(logger Logger) {
				logger.Info("info message")
			},
			expected: `level=info msg="info message"`,
		},
		{
			config: NewConfig("test"),
			logFunc: func(logger Logger) {
				logger.Infof("info message [%d]", 111)
			},
			expected: `level=info msg="info message [111]"`,
		},
		{
			config: NewConfig("test"),
			logFunc: func(logger Logger) {
				logger.Warn("warning message")
			},
			expected: `level=warning msg="warning message"`,
		},
		{
			config: NewConfig("test"),
			logFunc: func(logger Logger) {
				logger.Warnf("warning message [%d]", 111)
			},
			expected: `level=warning msg="warning message [111]"`,
		},
		{
			config: NewConfig("test"),
			logFunc: func(logger Logger) {
				logger.Error("error message")
			},
			expected: `level=error msg="error message"`,
		},
		{
			config: NewConfig("test"),
			logFunc: func(logger Logger) {
				logger.Errorf("error message [%d]", 111)
			},
			expected: `level=error msg="error message [111]"`,
		},

		{
			config: NewConfig("test"),
			logFunc: func(logger Logger) {
				assert.Panics(t, func() {
					logger.Panic("panic message")
				})
			},
			expected: `level=panic msg="panic message"`,
		},
		{
			config: NewConfig("test"),
			logFunc: func(logger Logger) {
				assert.Panics(t, func() {
					logger.Panicf("panic message [%d]", 111)
				})
			},
			expected: `level=panic msg="panic message [111]"`,
		},
		// not write log test patterns
		{
			config: NewConfig("test",
				WithLevel(Levels.Info),
			),
			logFunc: func(logger Logger) {
				logger.Debug("debug message")
			},
			expected: ``,
		},
		{
			config: NewConfig("test",
				WithLevel(Levels.Info),
			),
			logFunc: func(logger Logger) {
				logger.Debugf("debug message [%d]", 111)
			},
			expected: ``,
		},
		{
			config: NewConfig("test",
				WithLevel(Levels.Warn),
			),
			logFunc: func(logger Logger) {
				logger.Info("info message")
			},
			expected: ``,
		},
		{
			config: NewConfig("test",
				WithLevel(Levels.Warn),
			),
			logFunc: func(logger Logger) {
				logger.Infof("info message [%d]", 111)
			},
			expected: ``,
		},
		{
			config: NewConfig("test",
				WithLevel(Levels.Error),
			),
			logFunc: func(logger Logger) {
				logger.Warn("warning message")
			},
			expected: ``,
		},
		{
			config: NewConfig("test",
				WithLevel(Levels.Error),
			),
			logFunc: func(logger Logger) {
				logger.Warnf("warning message [%d]", 111)
			},
			expected: ``,
		},
		{
			config: NewConfig("test",
				WithLevel(Levels.Panic),
			),
			logFunc: func(logger Logger) {
				logger.Error("error message")
			},
			expected: ``,
		},
		{
			config: NewConfig("test",
				WithLevel(Levels.Panic),
			),
			logFunc: func(logger Logger) {
				logger.Errorf("error message [%d]", 111)
			},
			expected: ``,
		},
		{
			config: NewConfig("test",
				WithLevel(Levels.Fatal),
			),
			logFunc: func(logger Logger) {
				logger.Panic("panic message")
			},
			expected: ``,
		},
		{
			config: NewConfig("test",
				WithLevel(Levels.Fatal),
			),
			logFunc: func(logger Logger) {
				logger.Panicf("panic message [%d]", 111)
			},
			expected: ``,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			logger, buff := mockStandardLogger(test.config)
			test.logFunc(logger)
			if len(test.expected) > 0 {
				assert.Equal(t, test.expected+"\n", buff.String())
			} else {
				assert.Empty(t, buff.String())
			}
		})
	}
}

func TestNewStandardLogger(t *testing.T) {
	config := NewConfig("test",
		WithLevel(Levels.Debug),
	)
	actual := newStandardLogger(config)
	expected := &standardLogger{
		Logger: logrus.New(),
		config: config,
	}
	{
		expected.Level = logrus.DebugLevel
		expected.Out = os.Stderr
	}
	assert.Equal(t, expected, actual)
}

func mockStandardLogger(config *Config) (Logger, *bytes.Buffer) {
	buffer := new(bytes.Buffer)
	{
		config.out = buffer
	}
	logger := newStandardLogger(config).(*standardLogger)
	{
		logger.Formatter = &logrus.TextFormatter{
			DisableTimestamp: true,
		}
	}
	return logger, buffer
}
