package main

import "introducao/enderecos"

func main() {
	tipoEndereco := enderecos.TipoDeEndereco("Rodovia das Flores")
	println(tipoEndereco) // Deve imprimir "rua"
}
