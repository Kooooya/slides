package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan string, 10)
	go DoSomething(c)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		c <- "ビリケンさんに会いたい"
		SleepRandam(200)
	}
}

func DoSomething(c chan string) {
	timeout := time.After(500 * time.Millisecond)
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-timeout:
			fmt.Println("timed out")
			return
		}
	}
}

func SleepRandam(n int) {
	d := time.Duration(rand.Intn(n)) * time.Millisecond
	fmt.Printf("sleep %dms\n", d/time.Millisecond)
	time.Sleep(d)
}
