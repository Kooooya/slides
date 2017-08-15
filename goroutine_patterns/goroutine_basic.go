package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("Hello, 大阪!", i)
		}(i)
	}
	time.Sleep(time.Second)
}
