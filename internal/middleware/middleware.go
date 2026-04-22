package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

func Register(app *fiber.App) {
	app.Use(logger.New())
	app.Use(recover.New())

}
