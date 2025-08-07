package main

import "fmt"

func main() {
	// Outra forma de declarar uma váriavel:
	// var number int8 = 100

	number := int32(-100) // Números positivos e negativos
	fmt.Printf("%d", number)
	// Por padrão, é usado a conforme a arquitetura do
	// computador para determinar a quantidade de bits utilizado.
	// No caso, ficaria int64

	numbertwo := uint32(100) // Número 0 e positivos
	fmt.Printf("%d", numbertwo)

	// int32 = rune
	number3 := rune(123)
	fmt.Printf("%d", number3)

	// byte = uint8
	number4 := rune(1234)
	fmt.Printf("%d", number4)

	text := "contexto aleatório"
	fmt.Println(text)

}
