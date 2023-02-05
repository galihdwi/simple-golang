package main

import (
	"golangweb/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.IndexHandler)
	mux.HandleFunc("/product", handler.ProductHandler)

	log.Println("Starting server on :7000")

	err := http.ListenAndServe(":7000", mux)
	log.Fatal(err)
}
