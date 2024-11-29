package main

import "fmt"

func main() {
	// tamanho fixo e limitado
	var array1 [5]string = [5]string{"Robert Jr.", "Neymar", "Santos Drumont", "Joao Silva", "Jose Maria"}

	// tamanho dinamico, mais flexivel
	var slice []string = []string{"name", "age", "height"}

	slice = append(slice, "strong")

	slice2 := array1[1:3]

	// o slice aponta para um array, e reflete as alterações do array pai
	array1[1] = "alterar"

	fmt.Println(array1)
	fmt.Println(slice2)

	// array interno
	fmt.Println("---------------------")
	slice3 := make([]int64, 10, 15)

	/*
		a funcao make() vai criar inicialmente um array
		com 15 posicoes e retornar pra gente um slice
		com as 10 primeiras posicoes (no caso a cima).
	*/

	fmt.Println(slice3) // print length and capacity
	fmt.Printf("len: %d | cap: %d\n", len(slice3), cap(slice3))
}
