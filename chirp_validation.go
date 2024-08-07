package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type inputParameters struct {
	Body string `json:"body"`
}

type errorOutputValue struct {
	Error string `json:"error"`
}

type validOutputValue struct {
	Valid bool `json:"valid"`
}

func chirpValidationHandler(w http.ResponseWriter, r *http.Request) {
	defaultErrorResponse := errorOutputValue{
		Error: "Something went wrong",
	}
	log.Printf("Beginning chirp validation...")
	decoder := json.NewDecoder(r.Body)
	params := inputParameters{}
	err := decoder.Decode(&params)
	if err != nil {
		log.Printf("Error decoding parameters: %s", err)
		dat, err := json.Marshal(defaultErrorResponse)
		if err != nil {
			log.Printf("Error marshalling JSON: %s", err)
			w.WriteHeader(500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(dat)
		return
	}

	if len(params.Body) <= 140 {
		respBody := validOutputValue{
			Valid: true,
		}
		dat, err := json.Marshal(respBody)
		if err != nil {
			log.Printf("Error marshalling JSON: %s", err)
			dat, _ := json.Marshal(defaultErrorResponse)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(500)
			w.Write(dat)
			log.Printf("Completed chirp validation...")
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(dat)
		log.Printf("Completed chirp validation...")
		return
	} else {
		respBody := errorOutputValue{
			Error: "Chirp is too long",
		}
		dat, err := json.Marshal(respBody)
		if err != nil {
			log.Printf("Error marshalling JSON: %s", err)
			dat, _ := json.Marshal(defaultErrorResponse)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(500)
			w.Write(dat)
			log.Printf("Completed chirp validation...")
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		w.Write(dat)
		log.Printf("Completed chirp validation...")
		return
	}
}
