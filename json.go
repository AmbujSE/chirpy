package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

// This file is reserved for JSON helper functions

func cleanProfanity(text string) string {
	profaneWords := []string{"kerfuffle", "sharbert", "fornax"}
	words := strings.Split(text, " ")
	
	for i, word := range words {
		lowerWord := strings.ToLower(word)
		for _, profane := range profaneWords {
			if lowerWord == profane {
				words[i] = "****"
				break
			}
		}
	}
	
	return strings.Join(words, " ")
}

func (cfg *apiConfig) respondWithError(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	
	respBody := map[string]string{"error": msg}
	dat, err := json.Marshal(respBody)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		return
	}
	w.Write(dat)
}

func (cfg *apiConfig) respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		return
	}
	w.Write(dat)
}

func (cfg *apiConfig) handlerValidateChirp(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}
	
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		cfg.respondWithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}
	
	const maxChirpLength = 140
	if len(params.Body) > maxChirpLength {
		cfg.respondWithError(w, http.StatusBadRequest, "Chirp is too long")
		return
	}
	
	cleanedBody := cleanProfanity(params.Body)
	
	type returnVals struct {
		CleanedBody string `json:"cleaned_body"`
	}
	cfg.respondWithJSON(w, http.StatusOK, returnVals{CleanedBody: cleanedBody})
}