package handlers

import (
	"database/sql"
	"encoding/json"
	"go_web_server/internal/db"
	"log"
	"net/http"
	"strconv"
)

// User represents a user in the system
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}

// GetUsersHandler returns all users from the database
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	// Set content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Query all users from the database
	rows, err := db.DB.Query("SELECT id, name, email, full_name FROM users")
	if err != nil {
		log.Printf("Error querying users: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User

	// Iterate through the rows and build the users slice
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.FullName)
		if err != nil {
			log.Printf("Error scanning user row: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	// Check for errors from iterating over rows
	if err = rows.Err(); err != nil {
		log.Printf("Error iterating over user rows: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Encode users as JSON and send response
	if err := json.NewEncoder(w).Encode(users); err != nil {
		log.Printf("Error encoding users to JSON: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

// GetUserHandler returns a single user by ID from the database
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	// Set content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Extract user ID from URL path
	// For now, we'll use a simple approach - in a real app you'd use a router like gorilla/mux
	// This assumes the URL pattern is /api/users/{id}
	userID := r.URL.Path[len("/api/users/"):]
	if userID == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	// Query user by ID from the database
	var user User
	err := db.DB.QueryRow("SELECT id, name, email, full_name FROM users WHERE id = ?", userID).Scan(
		&user.ID, &user.Name, &user.Email, &user.FullName,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		log.Printf("Error querying user: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Encode user as JSON and send response
	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Printf("Error encoding user to JSON: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

// CreateUserHandler creates a new user in the database
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	// Set content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Parse request body
	var newUser User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Validate required fields
	if newUser.Name == "" || newUser.Email == "" {
		http.Error(w, "Name and email are required", http.StatusBadRequest)
		return
	}

	// Insert user into database
	result, err := db.DB.Exec("INSERT INTO users (name, email, full_name) VALUES (?, ?, ?)",
		newUser.Name, newUser.Email, newUser.FullName)
	if err != nil {
		log.Printf("Error inserting user: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Get the ID of the inserted user
	userID, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error getting last insert ID: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Set the ID and return the created user
	newUser.ID = int(userID)
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(newUser); err != nil {
		log.Printf("Error encoding user to JSON: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

// UpdateUserHandler updates an existing user in the database
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	// Set content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Extract user ID from URL path
	userID := r.URL.Path[len("/api/users/update/"):]
	if userID == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	// Parse request body
	var updatedUser User
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Validate required fields
	if updatedUser.Name == "" || updatedUser.Email == "" {
		http.Error(w, "Name and email are required", http.StatusBadRequest)
		return
	}

	// Update user in database
	result, err := db.DB.Exec("UPDATE users SET name = ?, email = ?, full_name = ? WHERE id = ?",
		updatedUser.Name, updatedUser.Email, updatedUser.FullName, userID)
	if err != nil {
		log.Printf("Error updating user: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Check if any rows were affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error getting rows affected: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Set the ID and return the updated user
	updatedUser.ID, _ = strconv.Atoi(userID)
	if err := json.NewEncoder(w).Encode(updatedUser); err != nil {
		log.Printf("Error encoding user to JSON: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
