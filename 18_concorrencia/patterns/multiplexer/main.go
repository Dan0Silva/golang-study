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
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return channel
}

func multiplexer(channelOne, channelTwo <-chan string) <-chan string {
	outputChannel := make(chan string)

	go func() {
		for {
			select {
			case message := <-channelOne:
				outputChannel <- message
			case message := <-channelTwo:
				outputChannel <- message
			}
		}
	}()

	return outputChannel
}

func main() {
	channel1 := write("ola mundo")
	channel2 := write("teste")

	channelFinal := multiplexer(channel1, channel2)

	for message := range channelFinal {
		fmt.Println(message)
	}

}
