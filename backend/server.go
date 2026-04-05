package main

import (
	"log"
	"task-manager/backend/config"
	"task-manager/backend/middleware"
	"task-manager/backend/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Create Fiber app
	app := fiber.New()

	// Enable CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173,http://localhost:5174",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))

	// Handle preflight
	app.Options("*", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	// Connect DB
	config.ConnectDB()

	// Public routes
	app.Post("/api/register", routes.Register)
	app.Post("/api/login", routes.Login)

	// Protected routes (require JWT)
	app.Get("/api/tasks", middleware.JWTMiddleware(), routes.GetTasks)
	app.Post("/api/tasks", middleware.JWTMiddleware(), routes.CreateTask)

	// Test route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Backend is running 🚀")
	})

	// Start server
	log.Fatal(app.Listen(":3000"))
}