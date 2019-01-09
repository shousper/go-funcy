package example

import "github.com/shousper/go-funcy/example/sub"

//go:generate funcy --generators SliceOf*,MapOf Bar
//go:generate funcy -k=Key Foo
//go:generate funcy -k E D

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
	// X this is just a test, c'mon
	X() sub.Trump
}
