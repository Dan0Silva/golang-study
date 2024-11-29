package main

import "fmt"

// funcao de retorno nomeado
func manyCalc(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2

	return
}

// funcao variatica
func variatic(numeros ...int) int {
	total := 0

	for _, n := range numeros {
		total += n
	}

	return total
}

// funcao recursiva
func fibonacci(position int) int {

	if position <= 2 {
		return 1
	}

	x := fibonacci(position-1) + fibonacci(position-2)
	return x
}

// funcao closure
func closure() func() {
	texto := "dentro de closure"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

func main() {
	func() {
		fmt.Println("funcao anonima")
	}()

	soma, subtracao := manyCalc(10, 5)
	fmt.Println(soma, subtracao)

	fmt.Println(variatic(1, 2, 3, 4, 5, 100))

	fmt.Println("fibonacci: ", fibonacci(8))

	fmt.Println("===================================")
	funcaoNova := closure()

	funcaoNova()
}
