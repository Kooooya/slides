package main

import (
	"fmt"
	"time"
)

func main() {
	c := TickTack()
	for i := 0; i < 10; i++ {
		t := <-c
		fmt.Println(t.msg, t.t)
	}
}

type Time struct {
	t   time.Time
	msg string
}

func TickTack() <-chan Time {
	c := make(chan Time)
	go func() {
		tick0, tick1 := time.Tick(100*time.Millisecond), time.Tick(200*time.Millisecond)
		for {
			select {
			case t := <-tick0:
				c <- Time{t, "interval 100ms"}
			case t := <-tick1:
				c <- Time{t, "interval 200ms"}
			}
		}
	}()
	return c
}
