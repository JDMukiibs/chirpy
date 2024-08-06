package main

import (
	"fmt"
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
	httpServerMux.HandleFunc("GET /api/metrics", apiCfg.middlewareMetrics)
	httpServerMux.HandleFunc("/api/reset", apiCfg.middlewareMetricsReset)
	server := &http.Server{
		Addr:    ":" + port,
		Handler: httpServerMux,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(server.ListenAndServe())
}

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileServerHits++
		w.Header().Add("Cache-Control", "no-cache")
		next.ServeHTTP(w, r)
	})
}

func (cfg *apiConfig) middlewareMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hits: %d", cfg.fileServerHits)))
}

func (cfg *apiConfig) middlewareMetricsReset(w http.ResponseWriter, r *http.Request) {
	cfg.fileServerHits = 0
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hits have been reset to %d", cfg.fileServerHits)))
}
