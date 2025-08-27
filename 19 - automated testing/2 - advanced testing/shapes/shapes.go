package formas

import "math"

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height // Method that calculates the area of the rectangle
}

type Shape interface {
	Area() float64 // Method that returns the area of the shape (return is a float64)
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius // Method that calculates the area of the Circle
}
