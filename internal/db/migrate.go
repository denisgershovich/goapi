package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func Migrate(db *sql.DB, migrationsDir string) error {
	files, err := filepath.Glob(filepath.Join(migrationsDir, "*.sql"))
	if err != nil {
		return err
	}

	for _, file := range files {
		content, err := os.ReadFile(file) // Use os.ReadFile instead of ioutil.ReadFile
		if err != nil {
			return err
		}

		fmt.Printf("Running migration: %s\n", filepath.Base(file))
		if _, err := db.Exec(string(content)); err != nil {
			return fmt.Errorf("migration %s failed: %w", file, err)
		}
	}

	log.Println("All migrations ran successfully")
	return nil
}
