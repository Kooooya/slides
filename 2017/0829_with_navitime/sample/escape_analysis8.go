package main

type S struct{}

func main() {
	v := S{}

	m := make(map[int]*S)
	m[0] = &v
}
