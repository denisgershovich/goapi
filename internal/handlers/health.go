package handlers

import (
	"go_web_server/internal/db"
	"net/http"
)

// HealthHandler checks DB connectivity and returns 200 OK if healthy
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	if err := db.DB.Ping(); err != nil {
		http.Error(w, "unhealthy", http.StatusServiceUnavailable)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
