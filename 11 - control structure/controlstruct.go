package main

import "fmt"

func main() {
	fmt.Println("Control Structures in Go!")

	number := -1

	if number > 15 {
		fmt.Println("Number is greater than 15")
	} else {
		fmt.Println("Number is not greater than 15")
	}

	if otherNumber := number; otherNumber > 0 {
		fmt.Println("Other number is positive")
	} else if otherNumber == 0 {
		fmt.Println("Other number is zero")
	} else {
		fmt.Println("Other number is negative")
	}

}
