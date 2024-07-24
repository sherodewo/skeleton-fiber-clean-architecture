package router

import (
	"github.com/gofiber/fiber/v2"
	"skeleton-fiber-clean-architecture/infrastructure/container"
)

func SetupRoutes(app *fiber.App) {
	// Set up dependency container
	cont, err := container.NewContainer()
	if err != nil {
		panic(err)
	}

	// Setup routes for different modules
	SetupUserRoutes(app, cont)
	SetupHistoryRoutes(app, cont)
}
