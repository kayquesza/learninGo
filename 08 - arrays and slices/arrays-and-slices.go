package main

import (
	"fmt"
)

func main() { 
	fmt.Println("Arrays and Slices")

	var arr1 [5]int 
	arr1[0] = 1
	fmt.Println(arr1)

	arr2 := [5]string{} // Inserir valores dos respectivos dentro: {}
	fmt.Println(arr2)

	// [...] Tamanho baseado na quantidade de valores inseridos
	// O maior ponto de "flexibilidade" de um [array]
	arr3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr3)

	slice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(slice)

	// Adicionar valor (ou valores) à um slice já existente
	slice = append(slice, 8, 9, 10)
	fmt.Println(slice)

	// Pegar apenas um intervalo determinado de um array ou slice
	slice2 := slice[3:6]
	fmt.Println(slice2)

	slice3 := arr1[0:1]
	fmt.Println(slice3)

	// Arrays Internos

	slice4 := make([]float32, 10, 15)
	fmt.Println(slice4)
	fmt.Println(len(slice4)) // lenght (Tamanho)
	fmt.Println(cap(slice4)) // capacidade

	slice4 = append(slice4, 12, 42, 13)
	fmt.Println("--------------")
	fmt.Println(slice4)
	fmt.Println(len(slice4)) 
	fmt.Println(cap(slice4))
}
