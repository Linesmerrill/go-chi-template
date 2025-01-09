package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Basic route
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the home page!"))
	})

	// Route with URL parameters
	r.Get("/user/{userID}", func(w http.ResponseWriter, r *http.Request) {
		userID := chi.URLParam(r, "userID")
		w.Write([]byte("User ID: " + userID))
	})

	// Route with query parameters
	r.Get("/search", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("q")
		w.Write([]byte("Search query: " + query))
	})

	// Route with POST method
	r.Post("/submit", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Form submitted!"))
	})

	// Route with middleware
	r.With(middleware.NoCache).Get("/nocache", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This response is not cached"))
	})

	http.ListenAndServe(":8080", r)
}
