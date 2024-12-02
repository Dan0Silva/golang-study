package address

import (
	"testing"
)

type testCase struct {
	address      string
	expectResult string
}

// comecar com a palavra Test com o primeiro T maiusculo
func TestAddressType(t *testing.T) {
	addressForTesting := []testCase{
		{"Rua Santos Dumont", "Rua"},
		{"Avenida Brasil", "Avenida"},
		{"Estrada dos alarmados", "Estrada"},
		{"Rodovia 180, kabreiro", "Rodovia"},
		{"Banana", "Invalid Type"},
	}

	for addressIndex, addressCase := range addressForTesting {
		resultTypeRecived := AddressType(addressCase.address)

		if resultTypeRecived != addressCase.expectResult {
			t.Errorf("type error on %d case: \n\t\t[%s, %s]", addressIndex, resultTypeRecived, addressCase.expectResult)
		}
	}
}
