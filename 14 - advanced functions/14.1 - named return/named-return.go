package main

// É chamado de retorno nomeado, pois o retorno da função é nomeado
func calculateSum(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func main() {
	soma, subtracao := calculateSum(10, 5)
	println("Soma:", soma)
	println("Subtração:", subtracao)

	resultadoSoma, resultadoSubtracao := calculateSum(20, 15)
	println("Resultado da Soma:", resultadoSoma)
	println("Resultado da Subtração:", resultadoSubtracao)
}
