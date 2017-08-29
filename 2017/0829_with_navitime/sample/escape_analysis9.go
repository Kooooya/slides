package main

type S struct{}

func main() {
	v := S{}

	m := []*S{}
	m[0] = &v
}
