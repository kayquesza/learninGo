package main

import "fmt"

// closure demonstrates the concept of closures in Go.
func closure() func() {
	text := "Inside the function: Closure"

	function := func() {
		fmt.Println(text)
	}
	return function
}

func main() {
	text := "Inside the function: Main"
	fmt.Println(text)

	newFunction := closure() 
	newFunction()
}
