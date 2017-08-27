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
	tickets := makeTickets(maxNum)
	for i := 0; ; i++ {
		i := i
		<-tickets
		go func() {
			Log(i)
			tickets <- true
		}()
		time.Sleep(time.Nanosecond)
	}
}
func makeTickets(n int) chan bool {
	var tickets = make(chan bool, n)
	for i := 0; i < n; i++ {
		tickets <- true
	}
	return tickets
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
