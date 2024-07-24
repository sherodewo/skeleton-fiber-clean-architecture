package controller

import (
	"github.com/gofiber/fiber/v2"
	"skeleton-fiber-clean-architecture/application/user/usecase"
	"strconv"
)

type UserController struct {
	UserUseCase *usecase.UserUseCase
}

func (uc *UserController) GetUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	user, err := uc.UserUseCase.GetUser(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(user)
}

func (uc *UserController) CreateUser(c *fiber.Ctx) error {
	type request struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := uc.UserUseCase.CreateUser(req.Name, req.Email, req.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusCreated)
}
