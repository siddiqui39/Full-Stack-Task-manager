
package routes

import (
	"context"
	"task-manager/backend/config"
	"task-manager/backend/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	//Create a new_User struct to hold the request data
	user := new(models.User)

	//Parse JSON body into user struct
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	//Hash the user's password before storing
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{

		})
	}
	user.Password = string(hashedPassword)

	//Insert the new user into the database
	_, err = config.DB.Exec(
		context.Background(),
		"INSERT INTO users (email, password) VALUES ($1, $2)",
		user.Email, user.Password,
	)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),

		})
	}

	//Return success massage
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully!",
	})
}