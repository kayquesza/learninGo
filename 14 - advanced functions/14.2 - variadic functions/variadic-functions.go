package main

import "fmt"

func sum(numbers ...int) int {
	total := 0
	for _, numbers := range numbers {
		total += numbers
	}
	return total
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(10, 20, 30))
	fmt.Println(sum(100, 200))
	fmt.Println(sum())

	totalSum := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("Total sum:", totalSum)

}
