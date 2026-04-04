package routes

import (
	"context"
	"task-manager/backend/config"
	"task-manager/backend/models"
	"task-manager/backend/utils"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	// Parse request body
	var data struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	// Get user from database (pgx query)
	var user models.User
	err := config.DB.QueryRow(
		context.Background(),
		"SELECT id, email, password FROM users WHERE email=$1",
		data.Email,
	).Scan(&user.ID, &user.Email, &user.Password)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}

	// Check password
	if !utils.CheckPassword(user.Password, data.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}

	// Generate JWT token
	token, err := utils.GenerateToken(uint(user.ID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not login",
		})
	}

	// Send response
	return c.JSON(fiber.Map{
		"message": "Login successful",
		"token":   token,
	})
}