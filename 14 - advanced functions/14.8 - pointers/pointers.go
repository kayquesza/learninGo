package main

import "fmt"

func invert(num int) int {
	return num * -1
}

func invertPointer(num *int) {
	*num = *num * -1
}

func main() {
	num := 5
	numInverted := invert(num)
	fmt.Println(numInverted)
	fmt.Println(num)

	newNum := 40
	fmt.Println(newNum)
	invertPointer(&newNum)
	fmt.Println(newNum)

}
