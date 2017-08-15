package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	for i := 0; i < 5; i++ {
		i := i
		wg.Add(1)
		go func() {
			fmt.Printf("hogehoge %d\n", i)
			wg.Done()
		}()
	}
	wg.Wait()
}
