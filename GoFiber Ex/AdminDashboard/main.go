package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// In-memory storage for simplicity
var users = []fiber.Map{
	{"ID": 1, "Name": "Alice"},
	{"ID": 2, "Name": "Bob"},
}
var orders = []fiber.Map{
	{"ID": 1, "Product": "Laptop", "Quantity": 1},
	{"ID": 2, "Product": "Phone", "Quantity": 2},
}
var stats = fiber.Map{
	"TotalUsers":  2,
	"TotalOrders": 2,
}

func main() {
	app := fiber.New()

	// Serve the HTML dashboard
	app.Get("/", dashboard)

	// API endpoints for fetching data
	app.Get("/api/users", listUsers)
	app.Get("/api/orders", listOrders)
	app.Get("/api/stats", getStats)

	// API endpoints for adding data
	app.Post("/api/users", addUser)
	app.Post("/api/orders", addOrder)
	app.Post("/api/stats", updateStats)

	log.Fatal(app.Listen(":3014"))
}

func dashboard(c *fiber.Ctx) error {
	return c.SendFile("./views/index.html")
}

func listUsers(c *fiber.Ctx) error {
	return c.JSON(users)
}

func listOrders(c *fiber.Ctx) error {
	return c.JSON(orders)
}

func getStats(c *fiber.Ctx) error {
	return c.JSON(stats)
}

func addUser(c *fiber.Ctx) error {
	user := new(struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	})

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	users = append(users, fiber.Map{"ID": user.ID, "Name": user.Name})
	stats["TotalUsers"] = len(users)
	return c.JSON(user)
}

func addOrder(c *fiber.Ctx) error {
	order := new(struct {
		ID       int    `json:"id"`
		Product  string `json:"product"`
		Quantity int    `json:"quantity"`
	})

	if err := c.BodyParser(order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	orders = append(orders, fiber.Map{"ID": order.ID, "Product": order.Product, "Quantity": order.Quantity})
	stats["TotalOrders"] = len(orders)
	return c.JSON(order)
}

func updateStats(c *fiber.Ctx) error {
	newStats := new(struct {
		TotalUsers  int `json:"totalUsers"`
		TotalOrders int `json:"totalOrders"`
	})

	if err := c.BodyParser(newStats); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	stats["TotalUsers"] = newStats.TotalUsers
	stats["TotalOrders"] = newStats.TotalOrders
	return c.JSON(stats)
}
