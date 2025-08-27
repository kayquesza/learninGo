package main

import (
	"fmt"
	"time"
)

func main() {
	// "<- chan string" only receives values from the channel
	channel := write("Hello, World!") // Calls the write function that returns a channel

	for i := 0; i < 10; i++ {
		fmt.Println(<-channel) // Receives values from the channel and prints
	}
}

func write(text string) <-chan string {
	channel := make(chan string) // Creates a channel to send strings

	go func() {
		for {
			channel <- fmt.Sprintf("Received value: %s", text) // Sends the formatted text to the channel
			time.Sleep(time.Millisecond * 500)
		}
	}()
	return channel // Returns the channel so others can receive the values

}
