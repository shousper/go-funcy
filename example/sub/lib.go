package sub

//go:generate funcy -path github.com/shousper/go-funcy/example/sub -type Apple -key-field Color -v

// Apple a delicious fruit!
type Apple struct {
	breed string
	Color [4]byte
}
