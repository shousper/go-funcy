package example

import "github.com/shousper/go-funcy/example/sub"

//go:generate funcy -type Bar -key-field ID -generators SliceOf*,MapOf
//go:generate funcy -type Foo -key-field Key
//go:generate funcy -type D -key-field E

// Bar similar to Foo
type Bar struct {
	ID string
}

// Foo a common name for junk
type Foo struct {
	Key            int
	StringField    string
	EmbeddedField  *Bar
	InterfaceField D
}

// D the fourth letter of the alphabet
type D interface {
	// E the fifth letter of the alphabet
	E() string
	// Apple the worlds first trillion dollar company
	Apple() *sub.Apple
}
