package logger

import (
	"bytes"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestSetupRootLogger(t *testing.T) {
	tests := []struct {
		name     string
		logFunc  func()
		expected string
	}{
		{
			logFunc: func() {
				Debug("debug message")
			},
			expected: `level=debug msg="debug message"`,
		},
		{
			logFunc: func() {
				Debugf("debug message [%d]", 111)
			},
			expected: `level=debug msg="debug message [111]"`,
		},
		{
			logFunc: func() {
				Info("info message")
			},
			expected: `level=info msg="info message"`,
		},
		{
			logFunc: func() {
				Infof("info message [%d]", 111)
			},
			expected: `level=info msg="info message [111]"`,
		},
		{
			logFunc: func() {
				Warn("warning message")
			},
			expected: `level=warning msg="warning message"`,
		},
		{
			logFunc: func() {
				Warnf("warning message [%d]", 111)
			},
			expected: `level=warning msg="warning message [111]"`,
		},
		{
			logFunc: func() {
				Error("error message")
			},
			expected: `level=error msg="error message"`,
		},
		{
			logFunc: func() {
				Errorf("error message [%d]", 111)
			},
			expected: `level=error msg="error message [111]"`,
		},

		{
			logFunc: func() {
				assert.Panics(t, func() {
					Panic("panic message")
				})
			},
			expected: `level=panic msg="panic message"`,
		},
		{
			logFunc: func() {
				assert.Panics(t, func() {
					Panicf("panic message [%d]", 111)
				})
			},
			expected: `level=panic msg="panic message [111]"`,
		},
	}
	buff := new(bytes.Buffer)
	{
		SetupRootLogger(NewConfig("test",
			WithLevel(Levels.Debug),
			WithMaxLevel(Levels.Fatal),
			WithOut(buff),
		))
		if logger, ok := rootLogger.(*standardLogger); ok {
			logger.Formatter = &logrus.TextFormatter{
				DisableTimestamp: true,
			}
		}
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			buff.Reset()
			test.logFunc()
			if len(test.expected) > 0 {
				assert.Equal(t, test.expected+"\n", buff.String())
			} else {
				assert.Empty(t, buff.String())
			}
		})
	}
}

func TestNewLogger(t *testing.T) {
	{
		configs := []*Config{
			NewConfig("test"),
		}
		actual := newLogger(configs...)
		expected := newStandardLogger(configs[0])
		assert.Equal(t, expected, actual)
	}
	{
		configs := []*Config{
			NewConfig("test1"),
			NewConfig("test2"),
		}
		actual := newLogger(configs...)
		expected := newMultiLogger(configs...)
		assert.Equal(t, expected, actual)
	}
	{
		configs := []*Config{}
		assert.PanicsWithValue(t, "configration not found", func() {
			newLogger(configs...)
		})
	}
}
