package routes

import (
	"context"
	"task-manager/backend/config"
	"task-manager/backend/models"

	"github.com/gofiber/fiber/v2"
)

// Create Task
func CreateTask(c *fiber.Ctx) error {
	var task models.Task
	if err := c.BodyParser(&task); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Get userId from JWT token
	userId := c.Locals("userId")
	if userId == nil {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}
	task.UserID = int(userId.(float64)) // JWT stores numbers as float64

	// Insert into DB
	query := `INSERT INTO tasks (title, user_id) VALUES ($1, $2) RETURNING id`
	err := config.DB.QueryRow(context.Background(), query, task.Title, task.UserID).Scan(&task.ID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create task"})
	}

	return c.JSON(task)
}

// Get Tasks (only for the logged-in user)
func GetTasks(c *fiber.Ctx) error {
	userId := c.Locals("userId")
	if userId == nil {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}

	rows, err := config.DB.Query(context.Background(),
		"SELECT id, title, user_id FROM tasks WHERE user_id=$1", int(userId.(float64)))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch tasks"})
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.UserID); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to scan task"})
		}
		tasks = append(tasks, task)
	}

	// Ensure tasks is never nil
	if tasks == nil {
		tasks = []models.Task{}
	}

	return c.JSON(tasks)
}
