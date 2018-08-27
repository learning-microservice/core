package logger

// Formatter is ...
type Formatter string

// Formatters is ...
var Formatters = struct {
	Text Formatter
	JSON Formatter
}{
	Text: Formatter("text"),
	JSON: Formatter("json"),
}
