package main

import (
	"fmt"
	"go_web_server/internal/db"
	"go_web_server/internal/handlers"
	middleware "go_web_server/internal/middlewares"
	"log"
	"net/http"
)

func main() {
	// Initialize the database.
	if err := db.Init(); err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}
	defer db.DB.Close()

	http.Handle("/api", middleware.Logger(http.HandlerFunc(handlers.HomeHandler)))

	port := ":8080"
	fmt.Println("Server is running on port", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
