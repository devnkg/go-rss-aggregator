package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// respondWithJson writes a JSON response with the specified HTTP status code
// It sets the Content-Type header to application/json and encodes the payload
func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Printf("Error encoding JSON response: %v", err)
	}
}

// respondWithError writes an error response as JSON with the specified status code
// It creates a standard error response format with an "error" field
func respondWithError(w http.ResponseWriter, code int, msg string) {
	type errResponse struct {
		Error string `json:"error"`
	}
	respondWithJson(w, code, errResponse{Error: msg})
}
