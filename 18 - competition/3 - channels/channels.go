package main

import (
	"fmt"
	"time"
)

func write(text string, channel chan string) {
	for i := 0; i < 5; i++ {
		channel <- text // Sends text to the channel
		time.Sleep(time.Second)
	}

	close(channel) // Closes the channel after sending all messages
}

func main() {
	channel := make(chan string)          // "Make" is the keyword to create a channel
	go write("Hello, World!", channel) // Starts the goroutine to write to the channel

	// for {
	//		message, open := <-channel // Waiting to receive the first value from the channel
	//		if !open {
	//			break // If the channel is closed, exit the loop
	//		}
	//		fmt.Println(message) // Prints the received message
	//	}

	for message := range channel {
		fmt.Println(message) // Prints the remaining messages from the channel
	}

	fmt.Println("End of program")
}
