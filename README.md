# go-funcy

Yet another "generic" code generation tool. Inspired by [typewriter](https://github.com/clipperhouse/typewriter), [go_generics](https://github.com/google/gvisor/tree/master/tools/go_generics) and [genny](https://github.com/cheekybits/genny).

## Usage

Basic CLI help:

```bash
funcy -h
```

```
Usage of funcy:
  -group-fields string
    	Name of fields to group by (comma delimited)
  -key-field string
    	Name of map key field (default "ID")
  -path string
    	Type import path
  -type string
    	Names of type to generate against
  -v	Verbose output
```

### Example

See the [example](./example), but given a `.go` with:

```go
package mypackage

// Foo a common name for junk
type Foo struct {
	Key            int
	StringField    string
	EmbeddedField  *Bar
	InterfaceField D
}
```

Add a `go:generate` like:

```go
//go:generate funcy -path to/my-package -type Foo -key-field Key
```

Or invoke manually, via `make`, whatever, and you'll get [this](./example/foo.funcy.go).
