package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		i := i
		go fmt.Printf("hogehoge %d\n", i)
	}
}
