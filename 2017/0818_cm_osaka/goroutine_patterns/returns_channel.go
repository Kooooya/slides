package main

import (
	"fmt"
	"time"
)

func main() {
	c := Say("だるまの串かつが食べたい！")
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %s\n", <-c)
	}
	fmt.Println("疲れた!")
}

func Say(msg string) <-chan string { // 送信専用のchannelを返す
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf(`%s %d`, msg, i)
			time.Sleep(100 * time.Millisecond)
		}
	}()
	return c
}
