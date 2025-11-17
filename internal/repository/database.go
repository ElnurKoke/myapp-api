package repository

import (
	"database/sql"
	l "elestial/internal/logger"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func InitializeDB(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}
	l.Logger.Info("Database connected - ", dataSourceName)
	err = runMigrations(db, "./migration")
	if err != nil {
		return nil, fmt.Errorf("failed to run migrations: %v", err)
	}
	l.Logger.Info("Migration complated")
	return db, nil
}

func runMigrations(db *sql.DB, migrationsDir string) error {
	files, err := os.ReadDir(migrationsDir)
	if err != nil {
		return fmt.Errorf("failed to read migrations directory: %v", err)
	}

	for _, file := range files {
		if !file.IsDir() {
			content, err := os.ReadFile(migrationsDir + "/" + file.Name())
			if err != nil {
				return fmt.Errorf("failed to read migration file %s: %v", file.Name(), err)
			}

			_, err = db.Exec(string(content))
			if err != nil {
				return fmt.Errorf("failed to execute migration file %s: %v", file.Name(), err)
			}
		}
	}

	return nil
}
