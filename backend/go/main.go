package main

import (
	"fmt"
	"net/http"
)

// HelloHandler responds with a "Hello, world!" message.
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

func main() {
	// Register the handler function for the "/" route.
	http.HandleFunc("/", HelloHandler)

	// Start the server on port 8080.
	fmt.Println("Server is listening on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
