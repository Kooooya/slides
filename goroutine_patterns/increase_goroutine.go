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
	for i := 0; ; i++ {
		go Log(i)
		time.Sleep(time.Nanosecond)
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
