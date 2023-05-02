package main

import (
	"time"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type App struct {
	fiber *fiber.App
}

func NewApp() App {
	server := fiber.New(fiber.Config{
		AppName:           "revil.dev API Server",
		CaseSensitive:     true,
		EnablePrintRoutes: true,
		Immutable:         true,
		JSONEncoder:       json.Marshal,
		JSONDecoder:       json.Unmarshal,
	})

	server.Use(compress.New(compress.Config{
		Level: compress.LevelBestCompression,
	}))

	server.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:10000",
		AllowMethods:     "POST,GET",
		AllowCredentials: true,
	}))

	server.Post("/gql", func(ctx *fiber.Ctx) error {
		body := new(GqlBody)

		if err := ctx.BodyParser(body); err != nil {
			return err
		}

		result := gql.Do(gql.Params{
			Context:        ctx.Context(),
			Schema:         schema,
			RequestString:  body.Query,
			VariableValues: body.Variables,
			OperationName:  body.Operation,
		})

		return ctx.JSON(result)
	})
	server.Static("/sandbox", "./public/sandbox.html")

	return App{
		fiber: server,
	}
}

func (app *App) Start() error {
	if err := app.fiber.Listen(":20000"); err != nil {
		return err
	}

	return nil
}

func (app *App) Shutdown() error {
	if err := app.fiber.ShutdownWithTimeout(2 * time.Second); err != nil {
		return err
	}

	return nil
}
