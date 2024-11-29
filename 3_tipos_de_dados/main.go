package main

import (
	"errors"
	"fmt"
)

func main() {
	/*
		int8, int16, int32, int64
		int = com base na arquitetura

		uint = valores positivos, apenas

		rune = suporta o mesmo que o int32
	*/
	var number byte = 255 // alias para uint8

	/*
		float32, float64
	*/
	var floating float32 = 123.45

	/*
		string = "para cortas de texto"
	*/
	var name string = "Danilo A. Silva"

	/*
		char e diferente em go

		ele pega o numero da tabela ascii equivalente ao caractere
	*/
	char := 'b'

	var erro error = errors.New("erro na bagaça")

	fmt.Println(name, number, floating, char, erro)
}
