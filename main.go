package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d", id)

	// Test
	io.WriteString(w, " - WriteString")
	fmt.Fprint(w, " - Fprint")
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("TODO: Display a form for creating a new snippet"))
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	w.Header().Set("X-Website", "afonso.dev")
	w.Header().Add("X-Website", "afonsodemori.com")
	// w.WriteHeader(201)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("TODO: Save a new snippet"))

	// Test
	headers := w.Header()

	for key, values := range headers {
		for _, value := range values {
			fmt.Printf("%s: %s\n", key, value)
		}
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home) // Restrict this route to exact matches on / only.
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Print("starting server on port :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
