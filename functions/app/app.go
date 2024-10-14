package app

import (
	"bufio"
	"embed"
	"encoding/json"
	"fmt"
	"github.com/Masterminds/sprig/v3"
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

//go:embed views/* components/*
var templates embed.FS

var debug = false

func SetupApp() *fiber.App {
	engine := html.NewFileSystem(http.FS(templates), ".gohtml")
	engine.Reload(debug)
	engine.Debug(debug)
	engine.AddFuncMap(sprig.FuncMap())
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

	checker := db.NewChecker(20 * time.Millisecond)

	var serverVersion = strconv.FormatInt(time.Now().Unix(), 10)

	// TODO: Close DB
	app.Get("/exercises/:exerciseId", func(ctx *fiber.Ctx) error {
		exerciseId := ctx.Params("exerciseId")
		exercise, found := exercises.ExercisesMap[exerciseId]
		if !found {
			return ctx.Status(404).SendString("Not found")
		}

		return ctx.Render("views/exercises", fiber.Map{
			"ID":        exercise.ID,
			"Title":     exercise.Title,
			"DebugMode": debug,
			"Expected": fiber.Map{
				"Headers": exercise.CorrectHeaders,
				"Rows":    exercise.Correct,
			},
			"Description":   exercise.Description,
			"ServerVersion": serverVersion,
			"Previous":      exercise.Previous,
			"Next":          exercise.Next,
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
			SetEventTriggerHeader(ctx, false, query, exerciseId)
			return ctx.Status(200).SendString(err.Error())
		}

		SetEventTriggerHeader(ctx, result.Success, query, exerciseId)
		return ctx.Render("components/output_table", fiber.Map{
			"Headers": result.Headers,
			"Rows":    result.Rows,
		})
	})

	if debug {
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
	}

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Render("views/index", fiber.Map{
			"Exercises":     exercises.Exercises,
			"DebugMode":     debug,
			"ServerVersion": serverVersion,
		})
	})

	app.Get("/:page", func(ctx *fiber.Ctx) error {
		page := ctx.Params("page")
		switch page {
		case "about":
			return ctx.Render("views/about", fiber.Map{
				"DebugMode":     debug,
				"ServerVersion": serverVersion,
			})
		default:
			return ctx.Render("views/not_found", fiber.Map{})
		}
	})

	return app
}

func SetEventTriggerHeader(ctx *fiber.Ctx, isSuccessful bool, query string, id string) {
	b, _ := json.MarshalIndent(fiber.Map{
		"query_evaluated": fiber.Map{
			"id":            id,
			"is_successful": isSuccessful,
			"query":         query,
		},
	}, "", "  ")
	ctx.Set("HX-Trigger", string(b))
}
