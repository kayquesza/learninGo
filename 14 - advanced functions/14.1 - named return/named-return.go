package main

// It's called named return because the function return is named
func calculateSum(n1, n2 int) (sum int, subtraction int) {
	sum = n1 + n2
	subtraction = n1 - n2
	return
}

func main() {
	sum, subtraction := calculateSum(10, 5)
	println("Sum:", sum)
	println("Subtraction:", subtraction)

	resultSum, resultSubtraction := calculateSum(20, 15)
	println("Sum Result:", resultSum)
	println("Subtraction Result:", resultSubtraction)
}
