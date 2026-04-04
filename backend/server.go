package main

import (
	"log"
	"task-manager/backend/config"
	"task-manager/backend/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Create Fiber app
	app := fiber.New(fiber.Config{
		// Optional: Add this to enable strict header handling
		// Prefork: false,
	})

	// Enable CORS middleware for React frontend
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173", // React dev server
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))

	//Handle prefLight options requests
	app.Options("*", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	// Connect to PostgreSQL
	config.ConnectDB()

	// Register the API register route
	app.Post("/api/register", routes.Register)

	// Test route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Backend is running 🚀")
	})

	// Start server on port 3000
	log.Fatal(app.Listen(":3000"))
}