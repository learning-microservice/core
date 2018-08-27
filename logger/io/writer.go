package io

// Closable is ...
type Closable interface {
	Close() error
}

// Flashable is ...
type Flashable interface {
	Flush() error
}
