package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

func main() {
	// Create new fiber instance
	app := fiber.New()

	// Add global middleware
	app.Use(recover.New())
	app.Use(logger.New())

	// Register routes
	app.Get("/", func(c fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(
			fiber.Map{
				"name":    "thalestmm-don",
				"version": "0.0.0",
			},
		)
	})

	// Serve the application
	if err := app.Listen(":2000"); err != nil {
		log.Fatal(err)
	}
}
