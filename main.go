package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(writer http.ResponseWriter, req *http.Request) {

	if req.URL.Path != "/" {
		http.NotFound(writer, req)
		return
	}

	writer.Write([]byte("Hello from Snippetbox"))
}

func snippetView(writer http.ResponseWriter, req *http.Request) {

	id, err := strconv.Atoi(req.URL.Query().Get("id"))
	if (err != nil) || (id < 1) {
		http.NotFound(writer, req)
		return
	}

	fmt.Fprintf(writer, "Display a specific snippet with ID %d...", id)
}

func snippetCreate(writer http.ResponseWriter, req *http.Request) {

	if req.Method != http.MethodPost {
		writer.Header().Set("Allow", http.MethodPost)
		http.Error(writer, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	writer.Write([]byte("Create a new snippet..."))
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
