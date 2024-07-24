package main

import (
	"flag"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq" // Pastikan untuk mengimpor driver PostgreSQL
	"log"
	"skeleton-fiber-clean-architecture/config"
	"skeleton-fiber-clean-architecture/infrastructure/database"
	"skeleton-fiber-clean-architecture/infrastructure/logger"
	"skeleton-fiber-clean-architecture/infrastructure/router"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %s", err)
	}

	// Inisialisasi logger
	logger.InitLogger(&cfg.Logger)

	migrateFlag := flag.String("migrate", "", "run database migrations (up or down)")
	flag.Parse()

	if *migrateFlag != "" {
		database.RunMigration(*migrateFlag)
		return
	}

	app := fiber.New()

	// Inisialisasi router
	router.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
