package main

type S struct{}

func main() {
	var x S
	_ = identity(x)
}

func identity(x S) S {
	return x
}
