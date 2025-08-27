// Unit Test

package enderecos

import "testing"

type testScenario struct {
	insertedAddress string
	expectedReturn  string
}

func TestAddressType(t *testing.T) {

	// Slice of test scenarios
	// Each scenario contains an address and the expected return
	testScenarios := []testScenario{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Estrada do Sol", "Estrada"},
		{"Rodovia dos Bandeirantes", "Rodovia"},
		{"Praça da Sé", "Desconhecido"},
		{"RUA DAS FLORES", "Rua"},
		{"avenida das flores", "Avenida"},
		{"estrada das flores", "Estrada"},
		{"rodovia das flores", "Rodovia"},
		{"", "Desconhecido"},
	}

	for _, scenario := range testScenarios {
		receivedAddressType := TipoDeEndereco(scenario.insertedAddress)
		if receivedAddressType != scenario.expectedReturn {
			t.Errorf("The received type is different from expected.\nExpected: %s\nReceived: %s", scenario.expectedReturn, receivedAddressType)
		} else {
			t.Logf("Test OK! Address: %s, Type: %s", scenario.insertedAddress, receivedAddressType)
		}
	}

}
