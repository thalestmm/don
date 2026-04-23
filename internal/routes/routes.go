package routes

import "github.com/gofiber/fiber/v3"

func Register(app *fiber.App) {
	app.Get("/", func(c fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(
			fiber.Map{
				"name":    "don",
				"version": "0.0.0",
			},
		)
	})
	app.Get("/health", func(c fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	api := app.Group("api")
	v1 := api.Group("v1")

	v1.Get("/", func(c fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(
			fiber.Map{"message": "api v1 root"},
		)
	})
}
