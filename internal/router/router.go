package router

import (
	"go_web_server/internal/handlers"
	middleware "go_web_server/internal/middlewares"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func New() http.Handler {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)

	// Basic routes
	r.Get("/api", handlers.HomeHandler)
	r.Get("/health", handlers.HealthHandler)

	// RESTful user routes
	r.Route("/api/users", func(r chi.Router) {
		r.Get("/", handlers.GetUsersHandler)       // GET /api/users
		r.Post("/", handlers.CreateUserHandler)    // POST /api/users
		r.Get("/{id}", handlers.GetUserHandler)    // GET /api/users/{id}
		r.Put("/{id}", handlers.UpdateUserHandler) // PUT /api/users/{id}
	})

	return r
}
