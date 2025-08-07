package main

import "fmt"

func main() {
	// Aritméticos
	soma := 1 + 2
	subtracao := 3 + 4
	divisao := 10 / 2
	multiplicacao := 7 * 8
	restoDaDivisao := 9 % 10

	fmt.Println(soma)
	fmt.Println(subtracao)
	fmt.Println(divisao)
	fmt.Println(multiplicacao)
	fmt.Println(restoDaDivisao)

	n1 := int32(100)
	n2 := int32(15)
	soma2 := n1 + n2
	fmt.Println(soma2)
	// Fim dos Aritméticos

	// Atribuição
	var variavel1 string = "String"
	variavel2 := "String"
	fmt.Println(variavel1, variavel2)
	// Fim dos Operadores de Atribuição

	// Operadores Relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)
	// Fim dos Operadores Relacionais

	// Operadores Lógicos
	// && (e)
	// || (ou)
	// ! (negação)
	// Fim dos Operadores Lógicos

	// Operadores Unários
	numero := 10
	numero++     // acrescenta 1 em "numero" (10 + 1)
	numero += 10 // acrescenta o valor mencionado ao valor "numero" (10 + 10)
	fmt.Println(numero)

	numero--     // subtrai 1 em "numero" (10 - 1)
	numero -= 20 // subtrai com valor especifico

	numero *= 3 // multiplica com o valor especifico
	numero /= 3 // divide com o valor especifico
	numero %= 3 // resto da divisao com valor especifico

	fmt.Println(numero)
	// Fim dos Operadores Unários

}
