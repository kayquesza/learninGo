package main

import "fmt"

func main() {
	var obj1 string = "Objeto 1" // Declaração explícita
	fmt.Println(obj1)

	obj2 := "Objeto 2" // Declaração implícita
	fmt.Println(obj2)

	var (
		obj3 string = "Objeto 3"
		obj4        = 54321
	)
	fmt.Println(obj3, obj4)

	obj5, obj6 := "Objeto 5", "12345"
	fmt.Println(obj5, obj6)

}
