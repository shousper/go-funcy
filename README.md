# go-funcy

Yet another "generic" code generation tool. Inspired by [typewriter](https://github.com/clipperhouse/typewriter), [go_generics](https://github.com/google/gvisor/tree/master/tools/go_generics) and [genny](https://github.com/cheekybits/genny).

## Usage

Basic CLI help:

```bash
funcy --help
```

```
funcy is a "generic" code generation tool

Usage:
  funcy [flags] TYPE

Flags:
  -g, --generators strings     list of generators to run
  -f, --group-fields strings   name of fields on type to group by
  -h, --help                   help for funcy
  -k, --key-field string       name of field on type to populate map key (default "ID")
  -p, --path string            Type import path, can be relative to GOPATH
  -v, --verbose                verbose mode
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
//go:generate funcy --path to/my-package -k Key Foo
```

Or invoke manually, via `make`, whatever, and you'll get [this](./example/foo.funcy.go).

## Generators

Maps (requires `--key-field` match):

- [MapOf](./fragments/map.go#L9)
- [MapOfKeys](./fragments/map.go#L19)
- [MapOfValues](./fragments/map.go#L43)
- [MapOfSelect](./fragments/map.go#L67)
- [MapOfGroupBys](./fragments/map.go#L100)

Slices:

- [SliceOf](./fragments/slice.go#L9)
- [SliceOfAsMap](./fragments/slice.go#L19) (requires `--key-field` match)
- [SliceOfGroupBys](./fragments/slice.go#L45)
