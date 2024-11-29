package main

import "fmt"

func main() {
	var number1 int64 = 10
	var number2 int64 = number1

	var number3 int64 = 100
	var number4 *int64 = &number3

	/*
		* = inicia um ponteiro, como em *int64
		& = retorna o ponteiro em si de algo, como em &number3
		* = também desreferencia a variável, para exibir
		    o valor que está naquele endereço
	*/

	fmt.Println(number1, number2)
	fmt.Println(number3, *number4)
}
