package main

import (
	"encoding/json"
	"net/http"
)

var PORT = ":9090"

type Book struct {
	ID     int
	Title  string
	Stock  int
	Author string
}

var books = []Book{
	{ID: 1, Title: "Golang Beginner", Stock: 10, Author: "siapa"},
	{ID: 2, Title: "Elixir Beginner", Stock: 35, Author: "kamukan"},
}

func main() {
	http.HandleFunc("/books", getBooks)

	http.ListenAndServe(PORT, nil)
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		json.NewEncoder(w).Encode(books)
		return
	}

	http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
}
