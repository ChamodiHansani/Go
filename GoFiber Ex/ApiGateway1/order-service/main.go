package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/api/v1/orders", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"orders": []string{"order1", "order2"}})
	})

	log.Fatal(app.Listen(":3044"))
}
