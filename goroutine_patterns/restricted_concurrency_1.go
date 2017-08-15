package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"runtime"
	"time"
)

func init() {
	PrintStats()
}

func main() {
	const maxNum = 50
	var c = make(chan int)

	for i := 0; i < maxNum; i++ {
		go func() {
			for {
				Log(<-c)
			}
		}()
	}

	for i := 0; ; i++ {
		c <- i
	}
}

func Log(i int) {
	fmt.Fprintf(ioutil.Discard, "%d", i)
	time.Sleep(5 * time.Second)
}

func PrintStats() {
	go func() {
		for {
			PrintMemStats()
			PrintGoroutineNum()
			time.Sleep(time.Second)
		}
	}()
}

func PrintMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	log.Printf("Mem = %vB\n", m.Sys)
}

func PrintGoroutineNum() {
	log.Printf("Goroutines: %d", runtime.NumGoroutine())
}
