package address

import (
	"strings"
)

func isValidTypeOfAddress(word string) bool {
	validTypes := []string{"rua", "avenida", "estrada", "rodovia"}
	isValid := false

	for _, addressType := range validTypes {
		if strings.ToLower(word) == addressType {
			isValid = true
		}
	}

	return isValid
}

// AddressType verifica se o endereço é de um tipo valido
func AddressType(address string) string {
	firstWord := strings.Split(address, " ")[0]
	isValid := isValidTypeOfAddress(firstWord)

	if isValid {
		return firstWord
	} else {
		return "Invalid Type"
	}
}
