package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Item structure
type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var items []Item

func main() {
	app := fiber.New()

	// Route to get all items
	app.Get("/items", func(c *fiber.Ctx) error {
		return c.JSON(items)
	})

	// Route to get a single item by ID
	app.Get("/items/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
		}

		for _, item := range items {
			if item.ID == id {
				return c.JSON(item)
			}
		}
		return c.Status(fiber.StatusNotFound).SendString("Item not found")
	})

	// Route to create a new item
	app.Post("/items", func(c *fiber.Ctx) error {
		item := new(Item)
		if err := c.BodyParser(item); err != nil {
			return err
		}
		items = append(items, *item)
		return c.JSON(item)
	})

	// Route to update an existing item
	app.Put("/items/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
		}

		updatedItem := new(Item)
		if err := c.BodyParser(updatedItem); err != nil {
			return err
		}

		for i, item := range items {
			if item.ID == id {
				items[i] = *updatedItem
				return c.JSON(updatedItem)
			}
		}
		return c.Status(fiber.StatusNotFound).SendString("Item not found")
	})

	// Route to delete an item
	app.Delete("/items/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
		}

		for i, item := range items {
			if item.ID == id {
				items = append(items[:i], items[i+1:]...)
				return c.SendString("Item deleted successfully")
			}
		}
		return c.Status(fiber.StatusNotFound).SendString("Item not found")
	})

	// Start the Fiber server
	app.Listen(":3004")
}
