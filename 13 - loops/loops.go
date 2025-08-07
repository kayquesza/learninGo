package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Implementing a Loop.")
	fmt.Println()

	for i := 1; i <= 3; i++ {
		fmt.Println("Iteration:", i)
		time.Sleep(time.Millisecond)
	}

	names := []string{"Alice", "Bob", "Charlie"}
	for index, name := range names {
		fmt.Printf("Index: %d, Name: %s\n", index, name)
		time.Sleep(time.Millisecond)
	}

	for index, letter := range "Hello!" {
		fmt.Printf("Index: %d, Letter: %s\n", index, string(letter))
		time.Sleep(time.Millisecond)
	}

	user := map[string]string{
		"name":     "Alice",
		"lastName": "Caval",
	}

	for key, value := range user {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}

}
