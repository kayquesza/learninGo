package formas

import "math"

type Retangle struct {
	width  float64
	height float64
}

type Circle struct {
	raio float64
}

func (r Retangle) Area() float64 {
	return r.width * r.height // Método que calcula a área do retângulo
}

type Forma interface {
	Area() float64 // Método que retorna a área da forma (o retorno é um float64)
}

func (c Circle) Area() float64 {
	return math.Pi * c.raio * c.raio // Método que calcula a área do Círculo
}
