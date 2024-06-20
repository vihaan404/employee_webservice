package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshaling Json %s", err)
	}
	w.WriteHeader(code)
	w.Write(dat)
}

func responseWithError(w http.ResponseWriter, code int, msg string) {
	type errorResponse struct {
		Error string `json:"error"`
	}
	if code > 499 {
		log.Printf("error 5XX %s", msg)
	}
	respondWithJson(w, code, errorResponse{
		Error: msg,
	})
}
