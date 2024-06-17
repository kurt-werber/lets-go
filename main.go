package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const (
	port = ":4000"
)

// handler function
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte(fmt.Sprintf("Display a snippet with id %d", id)))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a snippet"))
}

func main() {
	mux := http.NewServeMux()                             //routers are referred to as servemux in Go
	mux.HandleFunc("GET /{$}", home)                      //registers home function as handler for / URL pattern, {$} prevents it from acting as a catchall
	mux.HandleFunc("GET /snippet/view/{id}", snippetView) //wildcards can be used
	mux.HandleFunc("GET /snippet/create", snippetCreate)

	log.Printf("Starting server on %v", port)

	err := http.ListenAndServe(port, mux)
	log.Fatal(err)
}
