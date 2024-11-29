package main

import "fmt"

func main() {
	// declaracao explicita
	var nome string = "danilo"

	// declaracao implicita
	sobrenome := "araujo silva"

	// declaracao em grupo explicita
	var (
		variavel1 string = "test1"
		variavel2 string = "test2"
	)

	// declaracao em grupo implicita
	variavel3, variavel4 := "test3", "test4"

	// declaracao constante
	const constante1 string = "alouuu"

	fmt.Printf("nome: %s %s \n", nome, sobrenome)
	fmt.Printf("%s | %s | %s | %s \n",
		variavel1, variavel2, variavel3, variavel4)

	fmt.Println("contante: ", constante1)
}
