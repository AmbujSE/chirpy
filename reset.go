package main

import "net/http"

// Reset the fileserver hits counter back to 0
func (cfg *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {
	// 1. Safely set the atomic counter to 0
	cfg.fileserverHits.Store(0)

	// 2. Set the response headers
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	// 3. Write a clear confirmation message
	w.Write([]byte("Hits reset to 0"))
}