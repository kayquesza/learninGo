package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Retangle", func(t *testing.T) {
		r := Retangle{10, 20}
		areaEsperada := float64(200)
		areaRecebida := r.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("Área esperada: %0.2f, Área recebida: %0.2f", areaEsperada, areaRecebida)
		}

	})

	t.Run("Circle", func(t *testing.T) {
		c := Circle{10}
		areaEsperada := float64(math.Pi * 10 * 10) // Usando math.Pi para maior precisão
		areaRecebida := c.Area()
		if areaEsperada != areaRecebida {
			t.Fatalf("Área esperada: %0.2f, Área recebida: %0.2f", areaEsperada, areaRecebida)
		}
	})
}
