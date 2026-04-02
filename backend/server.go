
package main 

import (
	"log"
	"task-manager/backend/config"
	"task-manager/backend/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	//Connect to PostgreSQL
	config.ConnectDB()

	//Register the api7register route
	app.Post("/api/register", routes.Register)

	//Test route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Backend is running 🚀")
	})

	//Start server
	log.Fatal(app.Listen(":3000"))
}

