package main

import (
	"fmt"
	"time"
)

func write(text string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("I recive the value: %s", text)
			time.Sleep(time.Millisecond * 200)
		}
	}()

	return channel
}

func main() {
	channelMessages := write("hello world")

	for message := range channelMessages {
		fmt.Println(message)
	}
}
