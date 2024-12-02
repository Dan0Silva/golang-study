package main

import (
	"automate_test/address"
	"fmt"
)

func main() {
	text := "Rua dos bobos, numero zero"
	result := address.AddressType(text)

	fmt.Println(result)
}
