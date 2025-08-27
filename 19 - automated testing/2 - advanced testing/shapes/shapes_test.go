package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Rectangle", func(t *testing.T) {
		r := Rectangle{10, 20}
		expectedArea := float64(200)
		receivedArea := r.Area()

		if expectedArea != receivedArea {
			t.Fatalf("Expected area: %0.2f, Received area: %0.2f", expectedArea, receivedArea)
		}

	})

	t.Run("Circle", func(t *testing.T) {
		c := Circle{10}
		expectedArea := float64(math.Pi * 10 * 10) // Using math.Pi for greater precision
		receivedArea := c.Area()
		if expectedArea != receivedArea {
			t.Fatalf("Expected area: %0.2f, Received area: %0.2f", expectedArea, receivedArea)
		}
	})
}
