package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func main() {
	// Create a new Fiber instance
	app := fiber.New()

	// Configure the rate limiting middleware
	rateLimitConfig := limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			// Allow bypassing rate limiting for specific routes or conditions
			return false
		},
		Max:        5,               // Maximum number of requests
		Expiration: 1 * time.Minute, // Reset rate limit every minute
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP() // Use IP address as the key for rate limiting
		},
		LimitReached: func(c *fiber.Ctx) error {
			// Custom response when rate limit is reached
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": "Too many requests",
			})
		},
	}

	// Apply rate limiting globally
	app.Use(limiter.New(rateLimitConfig))

	// Define a simple endpoint
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the rate-limited API!")
	})

	// Define another endpoint for testing
	app.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("This is a test endpoint.")
	})

	// Start the Fiber server
	log.Fatal(app.Listen(":3009"))
}
