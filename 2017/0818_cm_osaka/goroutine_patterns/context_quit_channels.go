package main

import (
	"context"
	"fmt"
)

func main() {
	var (
		c           = make(chan int)
		ctx, cancel = context.WithCancel(context.Background())
	)
	const n = 50
	for i := 0; i < n; i++ {
		go LaunchWorker(ctx, c)
	}
	cancel()
	fmt.Println("That was refreshing")
}

func LaunchWorker(ctx context.Context, c chan int) {
	for {
		select {
		case <-c:
			// do something
		case <-ctx.Done():
			return
		}
	}
}
