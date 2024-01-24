package app

import (
	"bufio"
	"embed"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
	"github.com/magnuswahlstrand/sql-exercises/functions/db"
	"github.com/magnuswahlstrand/sql-exercises/functions/exercises"
	"net/http"
	"strconv"
	"time"
)

//go:embed views/*
var templates embed.FS

func SetupApp() *fiber.App {
	engine := html.NewFileSystem(http.FS(templates), ".gohtml")
	engine.Reload(true)
	engine.Debug(true)
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	//var allowedOrigins string
	//if os.Getenv("SST_STAGE") != "prod" {
	//	allowedOrigins = "http://localhost:63342/"
	//} else {
	//	allowedOrigins = "https://htmx.link"
	//}

	app.Use(cors.New(
		cors.Config{
			//AllowHeaders:     "hx-request,hx-current-url",
			AllowHeaders:     "*",
			AllowCredentials: false,
			//AllowOrigins:     allowedOrigins,
			AllowOrigins: "*",
			AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		},
	))
	app.Use(logger.New())

	checker := db.NewChecker()

	var serverVersion = strconv.FormatInt(time.Now().Unix(), 10)

	// TODO: Close DB
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Render("views/index", fiber.Map{
			"Expected": fiber.Map{
				"Headers": exercises.Exercises[0].CorrectHeaders,
				"Rows":    exercises.Exercises[0].Correct,
			},
			"ServerVersion": serverVersion,
		})
	})

	app.Get("/check/:exerciseId", func(ctx *fiber.Ctx) error {
		query := ctx.Query("query")
		exerciseId := ctx.Params("exerciseId")
		if query == "" {
			return ctx.Status(400).SendString("Missing query")
		}

		result, err := checker.Check(exerciseId, query)
		if err != nil {
			return ctx.Status(200).SendString(err.Error())
		}

		statusCode := 200
		if !result.Success {
			//statusCode = 400
			statusCode = 200
		}

		return ctx.Status(statusCode).Render("views/output_table", fiber.Map{
			"Headers": result.Headers,
			"Rows":    result.Rows,
		})
	})

	sseHandler := func(c *fiber.Ctx) error {
		version := c.Query("version")
		c.Set("Content-Type", "text/event-stream")
		c.Set("Cache-Control", "no-cache")
		c.Set("Connection", "keep-alive")
		c.Set("Transfer-Encoding", "chunked")

		c.Context().SetBodyStreamWriter(func(w *bufio.Writer) {
			if version != serverVersion {
				fmt.Fprintf(w, "event: trigger_reload\n")
				fmt.Fprintf(w, "data: \"\"\n\n")
				if err := w.Flush(); err != nil {
					fmt.Printf("Error while flushing: %v. Closing http connection.\n", err)
				}
			} else {
				fmt.Println("Same version! Don't trigger")
			}
			//msg := fmt.Sprintf("%d - the 2time is %v", i, time.Now())

			for {
				// TODO: lock here forever, instead?
				time.Sleep(1 * time.Second)
			}
		})

		return nil
	}

	app.Get("/sse", sseHandler)

	return app
}
