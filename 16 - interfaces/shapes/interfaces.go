package main

import "fmt"

type retangle struct {
	width  float64
	height float64
}

type circle struct {
	raio float64
}

func (r retangle) area() float64 {
	return r.width * r.height // Método que calcula a área do retângulo
}

type forma interface {
	area() float64 // Método que retorna a área da forma (o retorno é um float64)
}

func escreverArea(f forma) {
	fmt.Printf("A área da forma é: %0.2f\n", f.area())
}

func (c circle) area() float64 {
	return 3.14 * c.raio * c.raio // Método que calcula a área do círculo
}

func main() {
	fmt.Println("Interfaces in Go")
	r := retangle{10, 5}
	escreverArea(r)

	c := circle{10}
	escreverArea(c)
}
