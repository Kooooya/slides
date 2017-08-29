package main

type S struct {
	M *int
}

func main() {
	ref()
}

func ref() (z S) {
	var i = [5]int{1, 2, 3, 4, 5}
	_i := i[0]
	z.M = &_i
	return z
}
