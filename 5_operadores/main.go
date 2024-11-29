package main

import "fmt"

func main() {
	// aritimeticos
	soma := 1 + 1
	subtracao := 1 - 1
	multiplicacao := 1 * 1
	divisao := 1 / 1
	resto := 1 % 1

	fmt.Println(soma, subtracao, multiplicacao, divisao, resto)

	// relacionais
	fmt.Println(1 < 1)
	fmt.Println(1 <= 1)
	fmt.Println(1 == 1)
	fmt.Println(1 != 1)
	fmt.Println(1 >= 1)
	fmt.Println(1 > 1)

	// logicos
	fmt.Println(false && true)
	fmt.Println(false || true)
	fmt.Println(!false)

	// unarios
	number := 10
	number++
	number--
	number += 5
	number -= 5
	number *= 3

	fmt.Println(number)
}
