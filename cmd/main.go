package main

import (
	"fmt"
	"go_web_server/internal/db"
	"go_web_server/internal/handlers"
	middleware "go_web_server/internal/middlewares"
	"net/http"
)

func main() {
	db.Init()
	defer db.DB.Close()

	http.Handle("/api", middleware.Logger(http.HandlerFunc(handlers.HomeHandler)))

	port := ":8080"
	fmt.Println("Server is running on port", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
