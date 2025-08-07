package main

import "fmt"

func recuperarExecucao() {
	// O "r" armazena o resultado do "recover"
	if r := recover(); r != nil {
		fmt.Println("Recuperando de um panic:", r)
	}
}

func studentIsAproved(n1, n2 float64) bool {
	defer recuperarExecucao() // Chama a função de recuperação quando ocorrer um panic
	media := (n1 + n2) / 2

	if media > 7 {
		return true
	} else if media < 7 {
		return false
	}

	// O "panic" interrumpe a execução do programa e exibe uma mensagem de erro
	panic("A média não pode ser calculada")
}

func main() {
	fmt.Printf("O aluno está aprovado? %t\n", studentIsAproved(7, 7))
	fmt.Println()
}
