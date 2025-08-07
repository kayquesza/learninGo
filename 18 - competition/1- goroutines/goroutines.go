package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("Hello, World!") // goroutine 
	escrever("Go is awesome!")
}

func escrever(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
