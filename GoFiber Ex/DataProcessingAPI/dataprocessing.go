package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	// Use logger middleware
	app.Use(logger.New())

	// Route to receive data stream
	app.Post("/data", processData)

	log.Fatal(app.Listen(":3016"))
}

// Handler function to process incoming data
func processData(c *fiber.Ctx) error {
	type DataChunk struct {
		Data string `json:"data"`
	}

	// Parse incoming data
	chunk := new(DataChunk)
	if err := c.BodyParser(chunk); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Process data in a separate goroutine
	go handleData(chunk.Data)

	return c.JSON(fiber.Map{"status": "processing"})
}

// Function to handle data processing
func handleData(data string) {
	// Simulate data processing
	fmt.Printf("Processing data: %s\n", data)
	time.Sleep(2 * time.Second)
	fmt.Printf("Finished processing data: %s\n", data)
}
