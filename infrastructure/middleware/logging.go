package middleware

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"time"
)

// LoggingMiddleware logs detailed information about each request
func LoggingMiddleware(c *fiber.Ctx) error {
	// Record the start time
	start := time.Now()

	// Proceed to the next middleware or handler
	err := c.Next()

	// Record the end time
	stop := time.Now()

	// Log request details
	log.Printf("Method: %s, URL: %s, Status: %d, Latency: %s, Error: %v",
		c.Method(),
		c.OriginalURL(),
		c.Response().StatusCode(),
		stop.Sub(start),
		err,
	)

	return err
}
