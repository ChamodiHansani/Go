package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func main() {
	app := fiber.New()

	// User Service Route
	app.All("/api/v1/users/*", func(c *fiber.Ctx) error {
		return proxy.Do(c, "http://localhost:3043"+c.OriginalURL())
	})

	// Order Service Route
	app.All("/api/v1/orders/*", func(c *fiber.Ctx) error {
		return proxy.Do(c, "http://localhost:3044"+c.OriginalURL())
	})

	log.Fatal(app.Listen(":3045"))
}
