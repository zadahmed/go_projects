package main

import (
	"fmt"
	"log"
	"net/http"
)

const message string = "Hello world!"

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("Server Failed to Start %v", err)
	}
}
