package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := multiplex(write("Hello, World!"), write("Goodbye, World!"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-channel) // Receives messages from the multiplexed channel
	}
}

func multiplex(inputChannel1, inputChannel2 <-chan string) <-chan string {
	outputChannel := make(chan string)

	go func() {
		for {
			select {
			case message := <-inputChannel1:
				outputChannel <- fmt.Sprintf("Channel 1: %s", message)
			case message := <-inputChannel2:
				outputChannel <- fmt.Sprintf("Channel 2: %s", message)
			}
		}
	}()

	return outputChannel
}

func write(text string) <-chan string {
	channel := make(chan string) // Creates a channel to send strings

	go func() {
		for {
			channel <- fmt.Sprintf("Received value: %s", text)              // Sends the formatted text to the channel
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000))) // Simulates a random delay
			// The delay simulates variability in message production
		}
	}()

	return channel
}
