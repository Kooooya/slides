package main

import "fmt"

func main() {
	c := make(chan int)
	for i := 0; i < 50; i++ {
		go DoSomething(c)
	}
	fmt.Println("I'm so full")
}

func DoSomething(c chan int) {
	for {
		select {
		case _ = <-c:
			// do something
		}
	}
}
