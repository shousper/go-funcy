package sub

//go:generate funcy --path github.com/shousper/go-funcy/example/sub --key-field=Color -f breed -v Apple

// Apple a delicious fruit!
type Apple struct {
	breed string
	Color [4]byte
}

// Trump a special breed of idiot
type Trump struct {
	Name string
}