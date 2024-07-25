package main

import (
	"flag"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // Pastikan untuk mengimpor driver PostgreSQL
	"log"
	"skeleton-fiber-clean-architecture/config"
	"skeleton-fiber-clean-architecture/infrastructure/container"
	"skeleton-fiber-clean-architecture/infrastructure/database"
	"skeleton-fiber-clean-architecture/infrastructure/logger"
	"skeleton-fiber-clean-architecture/infrastructure/middleware"
	"skeleton-fiber-clean-architecture/infrastructure/router"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %s", err)
	}

	// Initialize OAuth2Config
	middleware.InitializeOAuth2Config()

	// Initialize Fiber logger
	logger.InitLogger(&cfg.Logger)

	// Handle database migrations
	migrateFlag := flag.String("migrate", "", "run database migrations (up or down)")
	flag.Parse()
	if *migrateFlag != "" {
		database.RunMigration(*migrateFlag)
		return
	}

	// Initialize the container
	cont, err := container.NewContainer()
	if err != nil {
		log.Fatalf("Error initializing container: %v", err)
	}

	// Initialize Fiber application
	app := fiber.New()

	// Use logging middleware for all routes
	app.Use(middleware.LoggingMiddleware)

	// Setup routes with custom router
	router.SetupRoutes(app, cont)

	// Add routes for OAuth2 login
	app.Get("/login", middleware.HandleLogin)
	app.Get("/callback", middleware.HandleGoogleCallback)

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
