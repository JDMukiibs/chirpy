package main

import (
	"log"
	"net/http"
)

func main() {
	const port = "8080"
	// This creates a new HTTP request multiplexer
	httpServerMux := http.NewServeMux()
	httpServerMux.Handle("/", http.FileServer(http.Dir(".")))
	server := http.Server{
		Addr:    "localhost:" + port,
		Handler: httpServerMux,
	}
	log.Printf("Serving on port: %s\n", port)
	log.Fatal(server.ListenAndServe())
}
