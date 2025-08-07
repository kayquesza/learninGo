package main

import (
	"fmt"
	"time"
)

func main() {
	// "<- chan string" apenas recebe valores do canal
	canal := escrever("Hello, World!") // Chama a função escrever que retorna um canal

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal) // Recebe valores do canal e imprime
	}
}

func escrever(text string) <-chan string {
	canal := make(chan string) // Cria um canal para enviar strings

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", text) // Envia o texto formatado para o canal
			time.Sleep(time.Millisecond * 500)
		}
	}()
	return canal // Retorna o canal para que outros possam receber os valores

}
