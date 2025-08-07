package main

import "fmt"

func main() {
	canal := make(chan string, 2) // Cria um canal com buffer de tamanho X (1)
	canal <- "Hello, World!"      // Envia uma mensagem para o canal
	canal <- "Hello, again!"      // Envia outra mensagem para o canal

	message := <-canal // Recebe a mensagem do canal
	fmt.Println(message)
	fmt.Println(<-canal) // Recebe e imprime a próxima mensagem do canal
}
