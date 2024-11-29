package main

import (
	"fmt"
	"time"
)

func write(s string) {
	for {
		fmt.Println(s)
		time.Sleep(time.Second)
	}
}

// Concorrencia != Paralelismo
func main() {
	go write("ola mundo") // goroutine
	write("mainu")
}
