package main

import (
	"log"

	"github.com/thalestmm/don/internal/middleware"
	"github.com/thalestmm/don/internal/routes"

	"github.com/gofiber/fiber/v3"
)

func main() {
	// Create new fiber instance
	app := fiber.New()

	// Add global middleware
	middleware.Register(app)

	// Register routes
	routes.Register(app)

	// Serve the application
	if err := app.Listen(":2000"); err != nil {
		log.Fatal(err)
	}
}
