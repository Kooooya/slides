package main

import "testing"

type S struct {
	M *int
}

func elm() (z S) {
	var i = [5]int{1, 2, 3, 4, 5}
	_i := i[0]
	z.M = &_i
	return z
}

func BenchmarkEscape0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		elm()
	}
}

func elm1() (z S) {
	var i = [5]int{1, 2, 3, 4, 5}
	z.M = &i[0]
	return z
}

func BenchmarkEscape1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		elm1()
	}
}
