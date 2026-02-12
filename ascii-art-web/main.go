package main

import (
	"encoding/json" 
	"html/template"
	"log"
	"net/http"
	"time" 

	"ascii-art-web/handlers"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":    "healthy",
		"service":   "ascii-art-web",
		"timestamp": time.Now().Format(time.RFC3339),
	})
}

func main() {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatal(err)
	}

	h := handlers.NewHandler(tmpl)

	http.HandleFunc("/", h.Home)
	http.HandleFunc("/ascii-art", h.Generate)
	http.HandleFunc("/health", healthHandler) 

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}