package enderecos

import "strings"

// AddressType receives a string representing an address and returns the address type
// (street, avenue, road, highway) or "unknown"
func TipoDeEndereco(endereco string) string {
	validTypes := []string{"rua", "avenida", "estrada", "rodovia"}

	lowercaseAddress := strings.ToLower(endereco)
	// "ToLower" converts the string to lowercase
	firstWordAddress := strings.Split(lowercaseAddress, " ")[0]
	// "Split" will create a slice of separations based on whitespace
	// Example: "Rua das Flores" -> ["Rua", "das", "Flores"]
	// Returns index [0] which is the first element of the slice ("Rua")

	validAddress := false
	for _, tipo := range validTypes {
		if tipo == firstWordAddress {
			validAddress = true
		}
	}

	if validAddress {
		return strings.Title(firstWordAddress)
	}

	return "Desconhecido"
}
