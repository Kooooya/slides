package main

var n = 0

func main() {
	for i := 0; i < 1000; i++ {
		go func() {
			n++
		}()
	}
}
