package main

import (
	"net/http"
)

// Reset the fileserver hits counter back to 0 and delete all users if in dev mode
func (cfg *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {
	// Check if platform is dev
	if cfg.platform != "dev" {
		cfg.respondWithError(w, http.StatusForbidden, "Forbidden")
		return
	}

	// 1. Safely set the atomic counter to 0
	cfg.fileserverHits.Store(0)

	// 2. Delete all users from the database
	err := cfg.db.DeleteAllUsers(r.Context())
	if err != nil {
		cfg.respondWithError(w, http.StatusInternalServerError, "Failed to reset database")
		return
	}

	// 3. Write a success response
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(""))
}
