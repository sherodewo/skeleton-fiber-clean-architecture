package middleware

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"time"
)

// LoggingMiddleware logs the details of each request
func LoggingMiddleware(c *fiber.Ctx) error {
	start := time.Now()

	// Proceed to the next middleware or handler
	err := c.Next()

	// Log the request details
	stop := time.Now()
	log.Printf("%s %s %v\n", c.Method(), c.OriginalURL(), stop.Sub(start))

	return err
}
