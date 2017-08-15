package main

import "sync"

var n = 0

func main() {
	m := new(sync.Mutex)
	for i := 0; i < 1000; i++ {
		go func() {
			m.Lock()
			n++
			m.Unlock()
		}()
	}
}
