package main

import "fmt"

func somar(n1, n2 int64) int64 {
	return n1 + n2
}

func calculos(n1, n2 int8) (int8, int8) {
	return (n1 + n2), (n1 * n2)
}

func main() {
	fmt.Println(somar(2, 2))

	f := func(txt string) string {
		fmt.Println(txt)
		return txt + "this"
	}

	test := f("ola mundo")
	fmt.Println(test)

	resultados1, _ := calculos(12, 6)
	fmt.Println(resultados1)
}
