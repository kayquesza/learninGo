package main

import "fmt"

// Function "Add" receives 2 parameters and their types (n1 int8, n2 int8),
// and returns a number (int8)
func add(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func mathematicalCalculations(n1, n2 int8) (int16, int8) {
	sum := n1 + n2
	subtraction := n1 - n2
	return int16(sum), subtraction
}

func main() {
	sum := add(20, 20)
	fmt.Printf("The sum of the numbers is: %d.", sum)

	funcf := func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	funcf("Improvise Text")

	resultSum, resultSubtraction := mathematicalCalculations(10, 20)
	fmt.Println(resultSum, resultSubtraction)

	resultSum2, _ := mathematicalCalculations(15, 12)
	fmt.Println(resultSum2)

	_, resultSum3 := mathematicalCalculations(15, 12)
	fmt.Println(resultSum3)

}
