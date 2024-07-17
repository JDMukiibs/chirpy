package main

import (
	"net/http"
)

func main() {
	httpServerMux := http.NewServeMux()
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: httpServerMux,
	}
	server.ListenAndServe()
}
