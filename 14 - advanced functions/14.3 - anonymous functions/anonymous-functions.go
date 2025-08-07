package main

import "fmt"

func main() {

	func() {
		fmt.Println("Função anônima executada!")
	}()

	func(text string) {
		fmt.Println(text)
	}("Função anônima com parâmetro!")

	// Função anônima com retorno armazenada
	retorno := func(text string) string {
		return fmt.Sprintf("Recebido - %s", text)
	}("Função anônima com parâmetro!")
	fmt.Println(retorno)

}
