package main

import (
	"log"
	"net/http"
)

// handler function
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a snippet"))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a snippet"))
}

func main() {
	mux := http.NewServeMux()    //routers are referred to as servemux in Go
	mux.HandleFunc("/{$}", home) //registers home function as handler for / URL pattern, {$} prevents it from acting as a catchall
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("sartomg server on: 4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
