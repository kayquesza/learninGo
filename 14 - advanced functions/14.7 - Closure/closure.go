package main

import "fmt"

// closure demonstrates the concept of closures in Go.
func closure() func() {
	text := "Dentro da função: Closure"

	funcao := func() {
		fmt.Println(text)
	}
	return funcao
}

func main() {
	text := "Dentro da função: Main"
	fmt.Println(text)

	newFunction := closure() 
	newFunction()
}
