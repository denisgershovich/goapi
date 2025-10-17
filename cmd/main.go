package main

import (
	"fmt"
	"go_web_server/internal/config"
	"go_web_server/internal/db"
	"go_web_server/internal/router"
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

	r := router.New()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Server is running on port", ":"+port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
