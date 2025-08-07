package main

import (
	"fmt"
	"time"
)

func escrever(text string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- text // Envia o texto para o canal
		time.Sleep(time.Second)
	}

	close(canal) // Fecha o canal após enviar todas as mensagens
}

func main() {
	canal := make(chan string)          // "Make" é a palavra-chave para criar um canal
	go escrever("Hello, World!", canal) // Inicia a goroutine para escrever no canal

	// for {
	//		mensagem, aberto := <-canal // Esperando receber o primeiro valor do canal
	//		if !aberto {
	//			break // Se o canal estiver fechado, sai do loop
	//		}
	//		fmt.Println(mensagem) // Imprime a mensagem recebida
	//	}

	for mensagem := range canal {
		fmt.Println(mensagem) // Imprime as mensagens restantes do canal
	}

	fmt.Println("Fim do programa")
}
