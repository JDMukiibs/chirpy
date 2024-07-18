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
	httpServerMux.Handle("/app/*", http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot))))
	httpServerMux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	server := &http.Server{
		Addr:    ":" + port,
		Handler: httpServerMux,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(server.ListenAndServe())
}
