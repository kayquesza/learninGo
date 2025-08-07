// Teste de Unidade

package enderecos

import "testing"

type cenarioDeTest struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {

	// Slice de cenários de teste
	// Cada cenário contém um endereço e o retorno esperado
	cenariosDeTest := []cenarioDeTest{
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

	for _, cenario := range cenariosDeTest {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido é diferente do esperado.\nEsperado: %s\nRecebido: %s", cenario.retornoEsperado, tipoDeEnderecoRecebido)
		} else {
			t.Logf("Teste OK! Endereço: %s, Tipo: %s", cenario.enderecoInserido, tipoDeEnderecoRecebido)
		}
	}

}
