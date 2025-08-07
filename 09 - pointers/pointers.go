package main

import "fmt"

func main() {
	fmt.Println("pointers")

	v1 := 10
	v2 := v1

	fmt.Println(v1, v2)

	v1++
	fmt.Println(v1, v2)

	// "Ponteiro" é uma REFERÊNCIA de memória
	var v3 int        // Valor de "v3"
	var ponteiro *int // Endereço de memória onde o valor de "v3" esta armazenado

	v3 = 15
	ponteiro = &v3
	fmt.Println(v3, ponteiro)
	fmt.Println(v3, *ponteiro) // A adição do * pega o que está no endereço

	fmt.Println()
	v3 = 20
	fmt.Println(v3, ponteiro)
	fmt.Println(v3, *ponteiro)

}
