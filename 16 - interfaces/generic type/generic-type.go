package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf) // Função genérica que aceita qualquer tipo de dado
}

func main() {
	generica("Hello, World!") // Passando uma string
	generica(42)              // Passando um inteiro
	generica(3.14)            // Passando um float
	generica(true)            // Passando um booleano

	mapa := map[interface{}]interface{}{
		1:             "String",
		float32(1000): true,
		"String":      "String",
	}

	fmt.Println(mapa) // Exibindo o mapa com chaves e valores de tipos variados
}
