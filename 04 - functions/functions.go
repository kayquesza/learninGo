package main

import "fmt"

// Função "Somar" recebe 2 parâmetros e seus tipos (n1 int8, n2 int8),
// e retorna um número (int8)
func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int16, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return int16(soma), subtracao
}

func main() {
	soma := somar(20, 20)
	fmt.Printf("A soma dos números é: %d.", soma)

	funcf := func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	funcf("Texto Improvisado")

	resultadoSoma, resultadoSubstracao := calculosMatematicos(10, 20)
	fmt.Println(resultadoSoma, resultadoSubstracao)

	resultadoSoma2, _ := calculosMatematicos(15, 12)
	fmt.Println(resultadoSoma2)

	_, resultadoSoma3 := calculosMatematicos(15, 12)
	fmt.Println(resultadoSoma3)

}
