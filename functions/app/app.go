package app

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/magnuswahlstrand/sql-exercises/functions/db"
)

func SetupApp() *fiber.App {
	var app *fiber.App
	app = fiber.New()
	app.Use(logger.New())

	checker := db.NewChecker()
	// TODO: Close DB
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello, World!")
	})
	app.Use("/", func(ctx *fiber.Ctx) error {
		fmt.Println("YEAH")
		fmt.Println(ctx.Queries())
		return ctx.Next()
	})

	app.Get("/check/:exerciseId", func(ctx *fiber.Ctx) error {
		query := ctx.Query("query")
		fmt.Println("orig", ctx.OriginalURL())
		fmt.Println("orig", ctx.BaseURL())
		exerciseId := ctx.Params("exerciseId")
		fmt.Println("XXX", query, exerciseId)
		if query == "" {
			return ctx.Status(400).SendString("Missing query")
		}

		correct, err := checker.Check(exerciseId, query)
		if err != nil {
			return ctx.Status(500).SendString(err.Error())
		}

		if !correct {
			return ctx.Status(400).SendString("Incorrect")
		}
		return ctx.Status(200).SendString("Correct")
	})

	return app
}
