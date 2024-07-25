package response

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

type Meta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

type DataResponse struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type ErrorResponse struct {
	Meta  Meta   `json:"meta"`
	Error string `json:"error"`
}

func logResponse(meta Meta, err error) {
	if err != nil {
		log.Printf("Status: %d, Message: %s, Error: %s", meta.Code, meta.Message, err.Error())
	} else {
		log.Printf("Status: %d, Message: %s", meta.Code, meta.Message)
	}
}

func Success(c *fiber.Ctx, data interface{}) error {
	meta := Meta{
		Code:    fiber.StatusOK,
		Message: "Success",
		Status:  "success",
	}
	logResponse(meta, nil)
	return c.Status(fiber.StatusOK).JSON(DataResponse{
		Meta: meta,
		Data: data,
	})
}

func BadRequest(c *fiber.Ctx, err error) error {
	meta := Meta{
		Code:    fiber.StatusBadRequest,
		Message: "Bad Request",
		Status:  "error",
	}
	logResponse(meta, err)
	return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
		Meta:  meta,
		Error: err.Error(),
	})
}

func Forbidden(c *fiber.Ctx, err error) error {
	meta := Meta{
		Code:    fiber.StatusForbidden,
		Message: "Forbidden",
		Status:  "error",
	}
	logResponse(meta, err)
	return c.Status(fiber.StatusForbidden).JSON(ErrorResponse{
		Meta:  meta,
		Error: err.Error(),
	})
}

func NotFound(c *fiber.Ctx, err error) error {
	meta := Meta{
		Code:    fiber.StatusNotFound,
		Message: "Not Found",
		Status:  "error",
	}
	logResponse(meta, err)
	return c.Status(fiber.StatusNotFound).JSON(ErrorResponse{
		Meta:  meta,
		Error: err.Error(),
	})
}

func InternalServerError(c *fiber.Ctx, err error) error {
	meta := Meta{
		Code:    fiber.StatusInternalServerError,
		Message: "Internal Server Error",
		Status:  "error",
	}
	logResponse(meta, err)
	return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
		Meta:  meta,
		Error: err.Error(),
	})
}

func CustomError(c *fiber.Ctx, code int, message string, err error) error {
	meta := Meta{
		Code:    code,
		Message: message,
		Status:  "error",
	}
	logResponse(meta, err)
	return c.Status(code).JSON(ErrorResponse{
		Meta:  meta,
		Error: err.Error(),
	})
}
