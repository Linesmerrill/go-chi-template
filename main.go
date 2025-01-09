package main

import (
	"net/http"
)

func main() {
	r := Routes()
	http.ListenAndServe(":8080", r)
}
