package router

import (
	"github.com/gofiber/fiber/v2"
	"skeleton-fiber-clean-architecture/infrastructure/container"
	"skeleton-fiber-clean-architecture/infrastructure/middleware"
)

// SetupRoutes initializes all the routes for the application
func SetupRoutes(app *fiber.App, cont *container.Container) {
	// Public routes
	app.Get("/login", middleware.HandleLogin)
	app.Get("/callback", middleware.HandleGoogleCallback)

	// Group API routes with authentication middleware
	api := app.Group("/api", middleware.AuthMiddleware)

	// Setup specific routes
	SetupUserRoutes(api, cont)
	SetupHistoryRoutes(api, cont)
}
