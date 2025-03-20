package main

import (
	"log"
	"net/http"
)

// Home handler to write a byte slice
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from Tinyblondeman!\n"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet...\n"))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating a new snippet...\n"))
}

func main() {
	// Init the server and register the home handler
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// Start server and listener + error handling
	log.Println("Starting server on port 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
