package main

import (
	"log"
	"net/http"
)

// Handlers
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from textShare"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet..."))
}

// Router
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Starting server on :3000")
	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}
