package main

import (
	"fmt"
	"net/http"
)

var PORT = ":5050"

func main() {
	http.HandleFunc("/", greet)
	http.ListenAndServe(PORT, nil)
}

func greet(w http.ResponseWriter, r *http.Request) {
	msg := "Wuihhhh"
	fmt.Fprintf(w, msg)
}
