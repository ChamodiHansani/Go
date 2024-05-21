package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Task represents a task with a title, description, and completion status.
type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

// In-memory storage for tasks.
var tasks = make(map[string]Task)

func main() {
	// Initialize a new Fiber app.
	app := fiber.New()

	// Define routes for CRUD operations.
	app.Post("/tasks", createTask)
	app.Get("/tasks/:id", getTask)
	app.Put("/tasks/:id", updateTask)
	app.Delete("/tasks/:id", deleteTask)
	app.Get("/tasks", getTasks)

	// Start the Fiber app on port 3000.
	log.Fatal(app.Listen(":3002"))
}

// createTask handles the creation of a new task.
func createTask(c *fiber.Ctx) error {
	task := new(Task)
	if err := c.BodyParser(task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	task.ID = uuid.NewString()
	tasks[task.ID] = *task

	return c.Status(fiber.StatusCreated).JSON(task)
}

// getTask retrieves a task by its ID.
func getTask(c *fiber.Ctx) error {
	id := c.Params("id")
	task, exists := tasks[id]
	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Task not found",
		})
	}
	return c.JSON(task)
}

// updateTask updates the details of an existing task.
func updateTask(c *fiber.Ctx) error {
	id := c.Params("id")
	task, exists := tasks[id]
	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Task not found",
		})
	}

	updateData := new(Task)
	if err := c.BodyParser(updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	task.Title = updateData.Title
	task.Description = updateData.Description
	task.Completed = updateData.Completed
	tasks[id] = task

	return c.JSON(task)
}

// deleteTask deletes a task by its ID.
func deleteTask(c *fiber.Ctx) error {
	id := c.Params("id")
	_, exists := tasks[id]
	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Task not found",
		})
	}

	delete(tasks, id)
	return c.SendStatus(fiber.StatusNoContent)
}

// getTasks retrieves all tasks.
func getTasks(c *fiber.Ctx) error {
	var taskList []Task
	for _, task := range tasks {
		taskList = append(taskList, task)
	}
	return c.JSON(taskList)
}
