package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"skeleton-fiber-clean-architecture/config"
	"skeleton-fiber-clean-architecture/infrastructure/logger"
)

func NewDBConnection() (*sql.DB, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.LogError(fmt.Errorf("Error loading config: %s", err))
		return nil, err
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.DBName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		logger.LogError(fmt.Errorf("Error opening database: %q", err))
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		logger.LogError(fmt.Errorf("Error connecting to the database: %q", err))
		return nil, err
	}

	logger.LogInfo("Successfully connected to the database!")
	return db, nil
}
