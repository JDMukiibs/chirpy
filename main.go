package main

import (
	"log"
	"net/http"
)

func main() {
	const filepathRoot = "."
	const port = "8080"

	// This creates a new HTTP request multiplexer
	httpServerMux := http.NewServeMux()
	httpServerMux.Handle("/", http.FileServer(http.Dir(filepathRoot)))
	server := &http.Server{
		Addr:    ":" + port,
		Handler: httpServerMux,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(server.ListenAndServe())
}
