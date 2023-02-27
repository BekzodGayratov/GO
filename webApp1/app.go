package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func check(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Health check")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/health", check)
	fmt.Println("Server starting")
	http.ListenAndServe(":3000", nil)
}
