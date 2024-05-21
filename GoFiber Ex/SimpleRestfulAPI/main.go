package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Using a map to store users with their IDs as keys for efficient access
var users = make(map[string]User)

func main() {
	app := fiber.New()

	// Define routes
	app.Post("/users", createUser)
	app.Get("/users/:id", getUser)
	app.Put("/users/:id", updateUser)
	app.Delete("/users/:id", deleteUser)
	app.Get("/users", getUsers)

	// Start the server
	log.Fatal(app.Listen(":3001"))
}

// Create a new user (Handler Functions)
func createUser(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	// Generate a unique ID for the user
	user.ID = uuid.NewString()
	// Add the user to the map
	users[user.ID] = *user

	return c.Status(fiber.StatusCreated).JSON(user)
}

// Get a user by ID
func getUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user, exists := users[id]
	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}
	return c.JSON(user)
}

// Update user details
func updateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user, exists := users[id]
	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	updateData := new(User)
	if err := c.BodyParser(updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	// Update user details
	user.Name = updateData.Name
	user.Email = updateData.Email
	users[id] = user

	return c.JSON(user)
}

// Delete a user by ID
func deleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	_, exists := users[id]
	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	// Delete the user from the map
	delete(users, id)
	return c.SendStatus(fiber.StatusNoContent)
}

// Get all users
func getUsers(c *fiber.Ctx) error {
	var userList []User
	for _, user := range users {
		userList = append(userList, user)
	}
	return c.JSON(userList)
}
