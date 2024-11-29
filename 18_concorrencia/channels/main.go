package main

import (
	"fmt"
	"time"
)

func write(s string, channel chan string) {
	for i := 0; i < 3; i++ {
		channel <- s // enviando a string s para o canal
		time.Sleep(time.Second)
	}

	close(channel)
}

func main() {
	channel := make(chan string)
	go write("ola mundo", channel)

	for {
		message, isOpen := <-channel // recebendo a string enviada no canal, pela linha 10

		if !isOpen {
			break
		}

		fmt.Println(message)
	}

	/*
		dá pra usar a seguinte forma para fazer isso funcionar também:

		for message := range channel {
			fmt.Println(message)
		}
	*/

}
