package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type user struct {
	Name  string
	Email string
}

func home(w http.ResponseWriter, r *http.Request) {

	user1 := user{
		"Kess",
		"kess@gmail.com",
	}

	templates.ExecuteTemplate(w, "home.html", user1)
}

var templates *template.Template

func main() {

	// Irá referenciar à arquivos que terminam com .hmtl
	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", home)

	fmt.Println("Running on port: 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))

}
