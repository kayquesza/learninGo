package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(2) // Adds two goroutines to wait

	go func() {
		write("Hello, World!")
		waitGroup.Done() // -1 from waitGroup.Add(2)
	}()

	go func() {
		write("Go is awesome!")
		waitGroup.Done() // -1 from waitGroup.Add(2)
	}()

	waitGroup.Wait() // Waits until all goroutines finish
}

func write(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
