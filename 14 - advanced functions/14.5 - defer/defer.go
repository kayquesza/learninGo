package main

import "fmt"

func function1() {
	fmt.Println("Function 1 executed")
}

func function2() {
	fmt.Println("Function 2 executed")
}

func studentApproved(number1, number2 float64) bool {
	average := (number1 + number2) / 2

	if average >= 6 {
		return true
	}
	return false
}

func main() {
	defer function1()
	function2()

	fmt.Println("Student approved:", studentApproved(7, 8))
	fmt.Println("Student approved:", studentApproved(5, 4))
}
