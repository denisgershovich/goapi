package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

// Connect opens a SQLite database at the given path
func Connect(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
