package logger

import (
	"io"
	"os"
)

// Config is ...
type Config struct {
	name      string
	level     Level
	maxLevel  Level
	formatter Formatter
	out       io.Writer
}

// ConfigOption is ...
type ConfigOption func(*Config)

// WithLevel is ...
func WithLevel(level Level) ConfigOption {
	return func(c *Config) {
		c.level = level
	}
}

// WithMaxLevel is ...
func WithMaxLevel(maxLevel Level) ConfigOption {
	return func(c *Config) {
		c.maxLevel = maxLevel
	}
}

// WithFormatter is ...
func WithFormatter(formatter Formatter) ConfigOption {
	return func(c *Config) {
		c.formatter = formatter
	}
}

// WithOut is ...
func WithOut(out io.Writer) ConfigOption {
	return func(c *Config) {
		c.out = out
	}
}

// NewConfig is ...
func NewConfig(name string, options ...ConfigOption) *Config {
	config := &Config{
		name:      name,
		level:     Levels.Debug,
		maxLevel:  Levels.Fatal,
		formatter: Formatters.Text,
		out:       os.Stderr,
	}
	for _, option := range options {
		option(config)
	}
	return config
}
