package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"os"
	"skeleton-fiber-clean-architecture/infrastructure/logger"
)

func RunMigration(direction string) {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		logger.LogError(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		logger.LogError(err)
	}

	migrationsDir := "migrations"

	switch direction {
	case "up":
		if err := goose.Up(db, migrationsDir); err != nil {
			logger.LogError(err)
		}
	case "down":
		if err := goose.Down(db, migrationsDir); err != nil {
			logger.LogError(err)
		}
	default:
		logger.LogError(err)
	}
	logger.LogError(err)
}
