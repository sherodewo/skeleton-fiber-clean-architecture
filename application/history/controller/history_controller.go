package controller

import (
	"github.com/gofiber/fiber/v2"
	"skeleton-fiber-clean-architecture/application/history/usecase"
	"strconv"
)

type HistoryController struct {
	HistoryUseCase *usecase.HistoryUseCase
}

func (hc *HistoryController) GetHistory(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	history, err := hc.HistoryUseCase.GetHistory(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(history)
}

func (hc *HistoryController) CreateHistory(c *fiber.Ctx) error {
	type request struct {
		ItemName  string `json:"item_name"`
		Quantity  int    `json:"quantity"`
		Action    string `json:"action"`
		CreatedBy int    `json:"created_by"`
	}
	var req request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := hc.HistoryUseCase.CreateHistory(req.ItemName, req.Quantity, req.Action, req.CreatedBy)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusCreated)
}
