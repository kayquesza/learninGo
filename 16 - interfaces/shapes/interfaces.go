package main

import "fmt"

type rectangle struct {
	width  float64
	height float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.width * r.height // Method that calculates the area of the rectangle
}

type shape interface {
	area() float64 // Method that returns the area of the shape (return is a float64)
}

func writeArea(f shape) {
	fmt.Printf("The area of the shape is: %0.2f\n", f.area())
}

func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius // Method that calculates the area of the circle
}

func main() {
	fmt.Println("Interfaces in Go")
	r := rectangle{10, 5}
	writeArea(r)

	c := circle{10}
	writeArea(c)
}
