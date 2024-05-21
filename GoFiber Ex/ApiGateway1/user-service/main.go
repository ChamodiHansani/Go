package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/api/v1/users", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"users": []string{"user1", "user2"}})
	})

	log.Fatal(app.Listen(":3043"))
}
