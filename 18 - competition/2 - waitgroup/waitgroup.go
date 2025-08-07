package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(2) // Adiciona duas goroutines à espera

	go func() {
		escrever("Hello, World!")
		waitGroup.Done() // -1 no waitGroup.Add(2)
	}()

	go func() {
		escrever("Go is awesome!")
		waitGroup.Done() // -1 no waitGroup.Add(2)
	}()

	waitGroup.Wait() // Aguarda até que todas as goroutines terminem
}

func escrever(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
