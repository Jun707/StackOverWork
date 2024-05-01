package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)

	log.Print("starting server on :8000")

	err := http.ListenAndServe(":8000", mux)
	log.Fatal(err)
}