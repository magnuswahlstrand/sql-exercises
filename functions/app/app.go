package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupApp() *fiber.App {
	var app *fiber.App
	app = fiber.New()
	app.Use(logger.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello, World!")
	})

	return app
}
