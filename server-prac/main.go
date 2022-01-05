package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 8080

	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/win", winHandler)

	fmt.Printf("Starting server on port %d", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]

	fmt.Fprintf(w, "Hello %s", path)
}

func winHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This URL is a winner!")
}
