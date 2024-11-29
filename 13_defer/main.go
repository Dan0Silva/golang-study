package main

import "fmt"

func funcao1() {
	fmt.Println("exec function 1")
}

func funcao2() {
	fmt.Println("exec function 2 ")
}

func main() {
	defer funcao1() // adiar a execucao até o ultimo momento possivel.
	funcao2()
}
