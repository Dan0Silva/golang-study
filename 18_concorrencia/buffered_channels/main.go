package main

import "fmt"

func main() {
	channel := make(chan string, 2) // adicionando um buffer no canal

	channel <- "ola mundo"

	message := <-channel
	fmt.Println(message)
}
