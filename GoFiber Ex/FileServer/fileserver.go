package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func main() {
	app := fiber.New()

	// Serve static files from the "public" directory
	app.Static("/", "./public")

	// Define a route to download a specific file
	app.Get("/download/:filename", func(c *fiber.Ctx) error {
		filename := c.Params("filename")
		return c.SendFile("./public/files/" + filename)
	})

	// Basic authentication middleware for access control on images
	app.Use("/images", basicauth.New(basicauth.Config{
		Users: map[string]string{
			"admin": "1234", // username: admin, password: 1234
		},
	}))

	// Serve the images directory with authentication
	app.Static("/images", "./public/images")

	log.Fatal(app.Listen(":3031"))
}
