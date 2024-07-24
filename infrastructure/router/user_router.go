package router

import (
	"github.com/gofiber/fiber/v2"
	userController "skeleton-fiber-clean-architecture/application/user/controller"
	"skeleton-fiber-clean-architecture/infrastructure/container"
)

func SetupUserRoutes(app *fiber.App, container *container.Container) {
	userCtrl := &userController.UserController{UserUseCase: container.UserUseCase}

	app.Get("/users/:id", userCtrl.GetUser)
	app.Post("/users", userCtrl.CreateUser)
}
