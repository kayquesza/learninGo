package enderecos

import "strings"

// TipoEndereco recebe uma string representando um endereço e retorna o tipo do endereço
// (rua, avenida, estrada, rodovia) ou "desconhecido"
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	enderecoMinusculo := strings.ToLower(endereco)
	// "ToLower" converte a string para minúsculas
	primeiraPalavraEndereco := strings.Split(enderecoMinusculo, " ")[0]
	// "Split" fará um slice de separações conforme o espaço em branco
	// Exemplo: "Rua das Flores" -> ["Rua", "das", "Flores"]
	// Retorna o indice [0] que é o primeiro elemento do slice ("Rua")

	enderecoValido := false
	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraEndereco {
			enderecoValido = true
		}
	}

	if enderecoValido {
		return strings.Title(primeiraPalavraEndereco)
	}

	return "Desconhecido"
}
