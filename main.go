package main

import (
	"log"

	"belajar-go-fiber/database"
	"belajar-go-fiber/handlers"
	"belajar-go-fiber/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables or defaults")
	}

	// Initialize database connection
	database.Connect()

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// Health check route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Task Management API is running!",
			"version": "1.0.0",
		})
	})

	// API routes
	api := app.Group("/api/v1")

	// Authentication routes (public)
	auth := api.Group("/auth")
	auth.Post("/register", handlers.Register)
	auth.Post("/login", handlers.Login)

	// Protected routes (require authentication)
	protected := api.Group("/", middleware.JWTProtected())

	// Task routes
	tasks := protected.Group("/tasks")
	tasks.Get("/", handlers.GetTasks)
	tasks.Get("/:id", handlers.GetTask)
	tasks.Post("/", handlers.CreateTask)
	tasks.Put("/:id", handlers.UpdateTask)
	tasks.Delete("/:id", handlers.DeleteTask)

	log.Println("Server starting on port 3000...")
	log.Fatal(app.Listen(":3000"))
}
