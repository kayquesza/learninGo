package main

import (
	"fmt"
	"time"
)

func main() {

	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Hello, Canal 1!"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Hello, Canal 2!"
		}
	}()

	for {
		select {
		case mensagemCanal1 := <-canal1:
			fmt.Println("Recebido do Canal 1:", mensagemCanal1)
		case mensagemCanal2 := <-canal2:
			fmt.Println("Recebido do Canal 2:", mensagemCanal2)
		}
	}
}
