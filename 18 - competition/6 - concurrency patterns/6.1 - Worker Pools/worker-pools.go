package main

import "fmt"

func fibonacci(position int) int {
	if position <= 1 {
		return position
	}

	return fibonacci(position-2) + fibonacci(position-1)
}

// "tools" é um canal que só recebe valores (tools <-chan int)
// "results" é um canal que só envia valores (results chan<- int)
// worker é uma função que processa ferramentas e envia resultados
func worker(tools <-chan int, results chan<- int) {
	for number := range tools {
		results <- fibonacci(number) // Calcula o Fibonacci do número recebido
	}
}

func main() {
	tools := make(chan int, 40)
	results := make(chan int, 40)

	go worker(tools, results) // Inicia o worker em uma goroutine
	go worker(tools, results)
	go worker(tools, results) // Inicia mais dois workers

	for i := 0; i < 41; i++ {
		tools <- i // Envia números para o canal de ferramentas
	}
	close(tools)

	for i := 0; i < 41; i++ {
		result := <-results                           // Recebe resultados do canal de resultados
		fmt.Printf("Fibonacci n%d°: %d\n", i, result) // Imprime o resultado do Fibonacci
	}

}
