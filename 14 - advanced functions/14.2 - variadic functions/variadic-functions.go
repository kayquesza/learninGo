package main

import "fmt"

func soma(numbers ...int) int {
	total := 0
	for _, numbers := range numbers {
		total += numbers
	}
	return total
}

func main() {
	fmt.Println(soma(1, 2, 3, 4, 5))
	fmt.Println(soma(10, 20, 30))
	fmt.Println(soma(100, 200))
	fmt.Println(soma())

	totalSoma := soma(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("Total da soma:", totalSoma)

}
