package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type user struct {
	Name string
	City string
	Age  uint8
}

var templates *template.Template

func home(w http.ResponseWriter, r *http.Request) {
	newUser := user{"Danilo A. Silva", "Unai", 22}

	templates.ExecuteTemplate(w, "index.html", newUser)
	// w.Write([]byte("cla"))
}

func main() {
	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", home)

	fmt.Println("Executando na porta 8080 ...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
