package db

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(dbURL string) error {
	cwd, _ := os.Getwd() // текущая рабочая директория
	migrationsPath := fmt.Sprintf("file://%s", filepath.Join(cwd, "migrations"))
	m, err := migrate.New(
		migrationsPath,
		dbURL,
	)
	if err != nil {
		return fmt.Errorf("migrate init error: %w", err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("migration failed: %w", err)
	}

	return nil
}
