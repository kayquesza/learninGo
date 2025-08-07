package auxiliar

import (
	"fmt"

	"github.com/badoux/checkmail"
)

// "Escrever" é uma função que escreve uma mensagem no console.
func Escrever() {
	fmt.Println("Escrevendo do pacote auxiliar")
	escrever2()
	// Por estar no mesmo pacote, a função "escrever2" pode ser acessada.

	err := checkmail.ValidateFormat("teste@gmail.com")
	fmt.Println(err) // <nil> se o email for válido
	if err != nil {
		fmt.Println(err)
	}

}
