package main

import (
	"fmt"
)

// Sintaxe de uma struct: type "name" struct {}
type usuario struct {
	name    string
	yearOld uint8
}

func main() {
	fmt.Println("Arquivo Struct")

	var u usuario
	u.name = "Kess"
	u.yearOld = 21
	fmt.Println(u)

	u2 := usuario{
		"Kessie",
		22,
	}
	fmt.Println(u2)

	u3 := usuario{
		yearOld: 12,
	}
	fmt.Println(u3)

}
