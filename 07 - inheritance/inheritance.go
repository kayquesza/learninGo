package main

import (
	"fmt"
)

type people struct {
	name     string
	lastName string
	yearOld  uint8
}

type student struct {
	people
	course     string
	university string
}

func main() {
	fmt.Println("O mais próximo de: Herança.")

	p1 := people{
		"Kess",
		"Souza",
		21,
	}
	fmt.Println(p1)

	e1 := student{
		p1,
		"Data Science",
		"Makenzie",
	}
	fmt.Println(e1)

}
