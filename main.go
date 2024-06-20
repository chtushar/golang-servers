package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	// Routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the home page!"))
	})

	r.Get("/api/profile", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the profile page!"))
	})

	// Start the server
	http.ListenAndServe(":8080", r)
}
