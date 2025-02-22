package main

import (
	"fmt"
	"go_web_server/internal/handlers"
	middleware "go_web_server/internal/middlewares"
	"net/http"
)

func main() {
	// Apply middleware to all routes
	http.Handle("/api", middleware.Logger(http.HandlerFunc(handlers.HomeHandler)))

	// Start the server
	port := ":8080"
	fmt.Println("Server is running on port", port)
	serverErr := http.ListenAndServe(port, nil)
	if serverErr != nil {
		fmt.Println("Error starting server:", serverErr)
	}
}
