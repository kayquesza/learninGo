package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	canal := multiplexar(escrever("Hello, World!"), escrever("Goodbye, World!"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal) // Recebe mensagens do canal multiplexado
	}
}

func multiplexar(canalDeEntrada1, canalDeEntrada2 <-chan string) <-chan string {
	canalDeSaida := make(chan string)

	go func() {
		for {
			select {
			case message := <-canalDeEntrada1:
				canalDeSaida <- fmt.Sprintf("Canal 1: %s", message)
			case message := <-canalDeEntrada2:
				canalDeSaida <- fmt.Sprintf("Canal 2: %s", message)
			}
		}
	}()

	return canalDeSaida
}

func escrever(text string) <-chan string {
	canal := make(chan string) // Cria um canal para enviar strings

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", text)              // Envia o texto formatado para o canal
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000))) // Simula um atraso aleatório
			// O atraso simula a variabilidade na produção de mensagens
		}
	}()

	return canal
}
