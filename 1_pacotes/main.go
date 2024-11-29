package main

import (
	"fmt"
	"test/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Ola mundo, diretamente da main")
	auxiliar.Escrever()

	error := checkmail.ValidateFormat("test@mail.com")
	fmt.Println(error)
}
