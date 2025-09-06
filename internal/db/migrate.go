package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func Migrate(db *sql.DB) error {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return fmt.Errorf("cannot get current file path")
	}

	migrationsDir := filepath.Join(filepath.Dir(filename), "../migrations")

	files, err := filepath.Glob(filepath.Join(migrationsDir, "*.sql"))
	if err != nil {
		return err
	}

	if len(files) == 0 {
		log.Println("No migrations found")
		return nil
	}

	for _, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			return fmt.Errorf("failed to read migration %s: %w", file, err)
		}

		fmt.Printf("Running migration: %s\n", filepath.Base(file))
		if _, err := db.Exec(string(content)); err != nil {
			return fmt.Errorf("migration %s failed: %w", file, err)
		}
	}

	log.Println("All migrations ran successfully")
	return nil
}
