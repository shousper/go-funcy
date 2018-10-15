package example

import "github.com/shousper/go-funcy/example/sub"

//go:generate funcy -path github.com/shousper/go-funcy/example -type Bar -key-field ID -v
//go:generate funcy -path github.com/shousper/go-funcy/example -type Foo -key-field Key -v
//go:generate funcy -path github.com/shousper/go-funcy/example -type D -key-field E -v

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
