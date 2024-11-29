package main

import (
	"fmt"
	"sync"
	"time"
)

func write(s string) {
	for i := 0; i < 10; i++ {
		fmt.Println(s)
		time.Sleep(time.Second)
	}
}

// Concorrencia != Paralelismo
func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2) // + 2

	go func() {
		write("ola mundo")
		waitGroup.Done() // - 1
	}()

	go func() {
		write("mainu")
		waitGroup.Done() // - 1
	}()

	waitGroup.Wait() // waitGroup == 0
}
