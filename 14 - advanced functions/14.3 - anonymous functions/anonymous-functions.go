package main

import "fmt"

func main() {

	func() {
		fmt.Println("Anonymous function executed!")
	}()

	func(text string) {
		fmt.Println(text)
	}("Anonymous function with parameter!")

	// Anonymous function with stored return
	returnValue := func(text string) string {
		return fmt.Sprintf("Received - %s", text)
	}("Anonymous function with parameter!")
	fmt.Println(returnValue)

}
