package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan string, 10)
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	go DoSomething(ctx, c)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		c <- "ビリケンさんに会いたい"
		SleepRandam(200)
	}
}

func DoSomething(ctx context.Context, c chan string) {
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-ctx.Done():
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
