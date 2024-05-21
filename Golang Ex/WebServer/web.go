package main

import (
	"net/http"
)

func main() {
	// Register a handler function for the root URL "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Serve the static HTML page
		http.ServeFile(w, r, "index.html")
	})

	// Start the web server on port 8080
	http.ListenAndServe(":8080", nil) // http://localhost:8080
}
