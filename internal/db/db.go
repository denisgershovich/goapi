package db

import (
	"database/sql"
	"log"
	"path/filepath"
)

var DB *sql.DB

func Init() {
	dbPath := filepath.Join("data", "app.db")
	migrationsDir := "migrations"

	database, err := Connect(dbPath)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	if err := Migrate(database, migrationsDir); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	DB = database
	log.Println("Database initialized and ready")
}
