package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pajri/url-shortener/handler"
)

func main() {
	http.HandleFunc("/", handler.HomeHandler)
	http.HandleFunc("/shorten", handler.ShortenHandler)
	http.HandleFunc("/r/", handler.RedirectHandler)

	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
