package router

import (
	"go_web_server/internal/handlers"
	middleware "go_web_server/internal/middlewares"
	"net/http"
)

// New builds and returns the application's HTTP handler with routes and middleware.
func New() http.Handler {
	mux := http.NewServeMux()

	// Routes
	mux.Handle("/api", middleware.Logger(http.HandlerFunc(handlers.HomeHandler)))
	mux.Handle("/health", middleware.Logger(http.HandlerFunc(handlers.HealthHandler)))

	return mux
}
