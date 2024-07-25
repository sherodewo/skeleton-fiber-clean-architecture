package router

import (
	"github.com/gofiber/fiber/v2"
	historyController "skeleton-fiber-clean-architecture/application/history/controller"
	"skeleton-fiber-clean-architecture/infrastructure/container"
)

func SetupHistoryRoutes(app fiber.Router, container *container.Container) {
	historyCtrl := &historyController.HistoryController{HistoryUseCase: container.HistoryUseCase}

	app.Get("/history", historyCtrl.GetHistory)
	app.Post("/history", historyCtrl.CreateHistory)
}
