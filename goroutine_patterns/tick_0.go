package main

import (
	"fmt"
	"time"
)

func main() {
	c := TickTack()
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
}

func TickTack() <-chan time.Time {
	c := make(chan time.Time)
	go func() {
		tick := time.Tick(100 * time.Millisecond)
		for now := range tick {
			c <- now
		}
	}()
	return c
}
