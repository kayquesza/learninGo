package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func users(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Loading users page..."))
}

func main() {

	// "HTTP" é um protoco de comunicação - Base da Comunicação WEB

	// Cliente (Faz a requisição) - Servidor (Processa requisição e envia a resposta)
	// Request - Response

	// Rotas
	// URI - Identificador do Recurso
	// Método - GET, POST, PUT, DELETE

	// Criando a rota "/home" que aparece a mensagem "Hello, World!"
	http.HandleFunc("/home", home)

	// Criando a rota "users"
	http.HandleFunc("/users", users)

	// Criando o servidor na porta 5000
	log.Fatal(http.ListenAndServe(":5000", nil))

}
