package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", homeHandler)
	r.With(AuthMiddleware).Get("/user/{userID}", userHandler)
	r.With(AuthMiddleware).Get("/search", searchHandler)
	r.With(AuthMiddleware).Post("/submit", submitHandler)
	r.With(middleware.NoCache, AuthMiddleware).Get("/nocache", noCacheHandler)

	return r
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the home page!"))
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	w.Write([]byte("User ID: " + userID))
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	w.Write([]byte("Search query: " + query))
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Form submitted!"))
}

func noCacheHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This response is not cached"))
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != "valid-token" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
