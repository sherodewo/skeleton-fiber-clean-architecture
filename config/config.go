package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	Database DatabaseConfig
	Logger   LoggerConfig
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

type LoggerConfig struct {
	LogFile    string
	MaxSize    int
	MaxAge     int
	MaxBackups int
	Compress   bool
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return nil, err
	}

	maxSize, err := strconv.Atoi(os.Getenv("LOG_MAX_SIZE"))
	if err != nil {
		return nil, err
	}

	maxAge, err := strconv.Atoi(os.Getenv("LOG_MAX_AGE"))
	if err != nil {
		return nil, err
	}

	maxBackups, err := strconv.Atoi(os.Getenv("LOG_MAX_BACKUPS"))
	if err != nil {
		return nil, err
	}

	config := &Config{
		Database: DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     port,
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   os.Getenv("DB_NAME"),
		},
		Logger: LoggerConfig{
			LogFile:    os.Getenv("LOG_FILE"),
			MaxSize:    maxSize,
			MaxAge:     maxAge,
			MaxBackups: maxBackups,
			Compress:   os.Getenv("LOG_COMPRESS") == "true",
		},
	}

	return config, nil
}
