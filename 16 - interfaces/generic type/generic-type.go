package main

import "fmt"

func generic(interf interface{}) {
	fmt.Println(interf) // Generic function that accepts any type of data
}

func main() {
	generic("Hello, World!") // Passing a string
	generic(42)              // Passing an integer
	generic(3.14)            // Passing a float
	generic(true)            // Passing a boolean

	myMap := map[interface{}]interface{}{
		1:             "String",
		float32(1000): true,
		"String":      "String",
	}

	fmt.Println(myMap) // Displaying the map with keys and values of varied types
}
