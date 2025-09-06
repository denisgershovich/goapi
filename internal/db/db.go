package db

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

var (
	DB   *sql.DB
	once sync.Once
)

func Init() error {
	var err error
	once.Do(func() {
		_, filename, _, ok := runtime.Caller(0)
		if !ok {
			err = errors.New("cannot get current file path")
			return
		}
		baseDir := filepath.Dir(filepath.Dir(filename))

		dbPath := os.Getenv("DB_PATH")
		if dbPath == "" {
			dataDir := filepath.Join(baseDir, "data")
			if err = os.MkdirAll(dataDir, 0755); err != nil {
				return
			}
			dbPath = filepath.Join(dataDir, "app.db")
		}

		database, connectErr := Connect(dbPath)
		if connectErr != nil {
			err = connectErr
			return
		}

		// Run migrations.
		if migrateErr := Migrate(database); migrateErr != nil {
			err = migrateErr
			return
		}

		DB = database
		log.Println("Database initialized and ready")
	})
	return err
}
