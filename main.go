package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
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

	r.Get("/api/profile/{username}", func(w http.ResponseWriter, r *http.Request) {
		username := chi.URLParam(r, "username")
		w.Write([]byte("Username: " + username))
	})

	// Start the server
	http.ListenAndServe(":8080", otelhttp.NewHandler(r, "server"))
}