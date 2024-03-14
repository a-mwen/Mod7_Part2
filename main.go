package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":3030", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// Getting the data from the request
	f, err := os.Open("./data.txt") // File to serve
	if err != nil {
		http.Error(w, "File not found", http.StatusInternalServerError)
		return
	}
	defer f.Close()
	io.Copy(w, f) // Sending it to the client
}
