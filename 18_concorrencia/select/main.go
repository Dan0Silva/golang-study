package main

import (
	"fmt"
	"time"
)

func main() {
	channelOne, ChannelTwo := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			channelOne <- "channel one"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			ChannelTwo <- "channel two"
		}
	}()

	for { // aplicando select para poupar tempo
		select {
		case messageChannelOne := <-channelOne:
			fmt.Println(messageChannelOne)

		case messageChannelTwo := <-ChannelTwo:
			fmt.Println(messageChannelTwo)
		}
	}
}
