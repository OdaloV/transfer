package main

import (
	"html/template"
	"log"
	"net/http"

	"ascii-art-web/handlers"
)

func main() {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatal(err)
	}

	h := handlers.NewHandler(tmpl)

	http.HandleFunc("/", h.Home)
	http.HandleFunc("/ascii-art", h.Generate)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}