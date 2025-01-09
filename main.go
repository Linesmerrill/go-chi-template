package main

import (
	"log"
	"net/http"

	"github.com/linesmerrill/go-chi-template/router"
)

func main() {
	r := router.Router()
	log.Println("Server is up and running on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
