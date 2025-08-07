package main

import "fmt"

func function1() {
	fmt.Println("Function 1 executed")
}

func function2() {
	fmt.Println("Function 2 executed")
}

func studentAproved(number1, number2 float64) bool {
	media := (number1 + number2) / 2

	if media >= 6 {
		return true
	}
	return false
}

func main() {
	defer function1()
	function2()

	fmt.Println("Student approved:", studentAproved(7, 8))
	fmt.Println("Student approved:", studentAproved(5, 4))
}
