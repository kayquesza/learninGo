package main

import (
	"fmt"
	"time"
)

func main() {
	go write("Hello, World!") // goroutine 
	write("Go is awesome!")
}

func write(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
