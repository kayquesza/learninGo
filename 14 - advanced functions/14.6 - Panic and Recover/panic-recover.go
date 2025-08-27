package main

import "fmt"

func recoverExecution() {
	// The "r" stores the result of "recover"
	if r := recover(); r != nil {
		fmt.Println("Recovering from panic:", r)
	}
}

func studentIsApproved(n1, n2 float64) bool {
	defer recoverExecution() // Calls the recovery function when panic occurs
	average := (n1 + n2) / 2

	if average > 7 {
		return true
	} else if average < 7 {
		return false
	}

	// "panic" interrupts program execution and displays an error message
	panic("The average cannot be calculated")
}

func main() {
	fmt.Printf("Is the student approved? %t\n", studentIsApproved(7, 7))
	fmt.Println()
}
