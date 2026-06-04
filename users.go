package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
}

func (cfg *apiConfig) handleUsers(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email string `json:"email"`
	}
	type response struct {
		User
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		// FIXED: Changed from StatusInternalServerError to StatusBadRequest
		cfg.respondWithError(w, http.StatusBadRequest, "Couldn't decode parameters")
		return
	}

	// FIXED: Added basic validation to ensure email isn't empty
	if params.Email == "" {
		cfg.respondWithError(w, http.StatusBadRequest, "Email is required")
		return
	}

	user, err := cfg.db.CreateUser(r.Context(), params.Email)
	if err != nil {
		cfg.respondWithError(w, http.StatusInternalServerError, "Couldn't create user")
		return
	}

	cfg.respondWithJSON(w, http.StatusCreated, response{
		User: User{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Email:     user.Email,
		},
	})
}
