package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLevels(t *testing.T) {
	assert.Equal(t, Levels, struct {
		Debug Level
		Info  Level
		Warn  Level
		Error Level
		Panic Level
		Fatal Level
	}{
		Debug: Level(1),
		Info:  Level(2),
		Warn:  Level(3),
		Error: Level(4),
		Panic: Level(5),
		Fatal: Level(6),
	})
}

func TestLevel_String(t *testing.T) {
	tests := []struct {
		name     string
		config   Level
		expected string
	}{
		{
			config:   Levels.Debug,
			expected: "debug",
		},
		{
			config:   Levels.Info,
			expected: "info",
		},
		{
			config:   Levels.Warn,
			expected: "warn",
		},
		{
			config:   Levels.Error,
			expected: "error",
		},
		{
			config:   Levels.Panic,
			expected: "panic",
		},
		{
			config:   Levels.Fatal,
			expected: "fatal",
		},
		{
			config:   Level(128),
			expected: "unknown",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.config.String())
		})
	}
}

func TestParseLevel(t *testing.T) {
	tests := []struct {
		name     string
		actual   string
		expected Level
	}{
		{
			actual:   "debug",
			expected: Levels.Debug,
		},
		{
			actual:   "info",
			expected: Levels.Info,
		},
		{
			actual:   "warn",
			expected: Levels.Warn,
		},
		{
			actual:   "error",
			expected: Levels.Error,
		},
		{
			actual:   "panic",
			expected: Levels.Panic,
		},
		{
			actual:   "fatal",
			expected: Levels.Fatal,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, ParseLevel(test.actual))
		})
	}
}

func TestParseLevel_InvalidLogLevel(t *testing.T) {
	assert.PanicsWithValue(t, "invalid log level [unknown]", func() {
		ParseLevel("unknown")
	})
}
