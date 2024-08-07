package main

import (
	"log"
	"net/http"
)

type apiConfig struct {
	fileServerHits int
}

func main() {
	log.SetPrefix("chirpy > ")
	const filepathRoot = "."
	const port = "8080"
	apiCfg := &apiConfig{}

	// This creates a new HTTP request multiplexer
	httpServerMux := http.NewServeMux()
	handler := http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot)))
	httpServerMux.Handle("/app/*", apiCfg.middlewareMetricsInc(handler))

	// To limit methods that are used for certain endpoints can follow
	// a pattern of [METHOD ][HOST][/PATH]
	httpServerMux.HandleFunc("GET /api/healthz", handlerReadiness)
	httpServerMux.HandleFunc("GET /admin/metrics", apiCfg.middlewareMetrics)
	httpServerMux.HandleFunc("/api/reset", apiCfg.middlewareMetricsReset)
	httpServerMux.HandleFunc("POST /api/validate_chirp", chirpValidationHandler)
	server := &http.Server{
		Addr:    ":" + port,
		Handler: httpServerMux,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(server.ListenAndServe())
}
