package main

import (
	"fmt"
	"go_web_server/internal/config"
	"go_web_server/internal/db"
	"go_web_server/internal/handlers"
	middleware "go_web_server/internal/middlewares"
	"log"
	"net/http"
	"os"
)

func main() {
	// Load env vars from .env if present.
	if err := config.Load(); err != nil {
		log.Fatalf("failed to load .env: %v", err)
	}

	// Initialize the database.
	if err := db.Init(); err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}
	defer db.DB.Close()

	http.Handle("/api", middleware.Logger(http.HandlerFunc(handlers.HomeHandler)))
	http.Handle("/health", middleware.Logger(http.HandlerFunc(handlers.HealthHandler)))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Server is running on port", ":"+port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
