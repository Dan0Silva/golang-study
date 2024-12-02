package address

import (
	"testing"
)

// comecar com a palavra Test com o primeiro T maiusculo
func TestAddressType(t *testing.T) {
	addressForTesting := "Avenida mano brow"
	expectedResult := "Avenida"
	result := AddressType(addressForTesting)

	if result != expectedResult {
		t.Error("Tipo recebido é diferente do esperado")
	}
}
