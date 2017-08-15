package main

import "fmt"

func main() {
	var (
		c    = make(chan int)
		quit = make(chan bool)
	)

	const n = 50
	for i := 0; i < n; i++ {
		go LaunchWorker(c, quit)
	}

	for i := 0; i < n; i++ {
		quit <- true
	}

	fmt.Println("That was refreshing")
}

func LaunchWorker(c chan int, quit chan bool) {
	for {
		select {
		case <-c:
			// do something
		case <-quit:
			return
		}
	}
}
