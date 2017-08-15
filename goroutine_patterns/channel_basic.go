package main

import "fmt"

func main() {
	ch := make(chan int)
	const n = 10

	for i := 0; i < n; i++ {
		go func(i int) {
			ch <- i
		}(i)
	}

	for i := 0; i < n; i++ {
		fmt.Println("ビリケンさんに会いたい", <-ch)
	}
}
