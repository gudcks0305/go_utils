package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Create a simple web server
	// The server will listen on port 8080
	// and respond with "Hello, world" to all requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, world")
	})

	http.ListenAndServe(":8080", nil)
}
