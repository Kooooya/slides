package main

import (
	"fmt"
	"time"
)

func main() {
	daruma := Say("だるまの串かつが食べたい！")
	tengu := Say("てんぐの串かつが食べたい！")
	c := FanIn(daruma, tengu)
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

func FanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() { for { c <- <-input1 } }()
	go func() { for { c <- <-input2 } }()
	return c
}
