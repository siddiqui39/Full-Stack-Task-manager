package routes

import (
	"task-manager/backend/config"
	"task-manager/backend/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v5"
)

// Register a new user
func Register(c *fiber.Ctx) error {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to hash password",
		})
	}

	// Insert user into DB
	query := `INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id`
	var userID int
	err = config.DB.QueryRow(c.Context(), query, input.Email, string(hashedPassword)).Scan(&userID)
	if err != nil {
		// Generic error response for unique email or other insertion errors
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "Email already exists or another error occurred",
		})
	}

	return c.JSON(fiber.Map{
		"id":    userID,
		"email": input.Email,
	})
}

// Login existing user
func Login(c *fiber.Ctx) error {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	// Get user from DB
	var user models.User
	err := config.DB.QueryRow(c.Context(),
		"SELECT id, email, password FROM users WHERE email=$1",
		input.Email,
	).Scan(&user.ID, &user.Email, &user.Password)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}

	// Compare hashed password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}

	// Create JWT token using config.JwtSecret
	claims := jwt.MapClaims{
		"userId": user.ID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(config.JwtSecret)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not login",
		})
	}

	return c.JSON(fiber.Map{
		"token": t,
	})
}