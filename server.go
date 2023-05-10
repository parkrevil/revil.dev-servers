package main

import (
	"log"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/storage/redis/v2"
)

type Server struct {
	config Config
	server *fiber.App
}

func NewServer(config Config) Server {
	config.print()
	server := fiber.New(fiber.Config{
		AppName:       "revil.dev",
		Immutable:     true,
		CaseSensitive: true,
		StrictRouting: true,
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
	})
	server.Use(recover.New())
	server.Use(cors.New(cors.Config{
		AllowOrigins:     os.Getenv("SERVER_CORS_ORIGINS"),
		AllowHeaders:     "Referer, Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "POST,GET",
		AllowCredentials: false,
	}))
	server.Use(compress.New(compress.Config{
		Level: compress.LevelBestCompression,
	}))
	server.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))

	/*
		TODO: Add custom logging middleware with Zap
	*/

	if config.env == Production {
		redisPort, err := strconv.Atoi(os.Getenv("REDIS_PORT"))
		if err != nil {
			log.Fatalf("Invalid redis port: %v", err)
		}

		server.Use(limiter.New(limiter.Config{
			Max:               10,
			Expiration:        10 * time.Second,
			LimiterMiddleware: limiter.SlidingWindow{},
			Storage: redis.New(redis.Config{
				Host:      os.Getenv("REDIS_HOST"),
				Port:      redisPort,
				Password:  os.Getenv("REDIS_PASSWORD"),
				Database:  0,
				Reset:     false,
				TLSConfig: nil,
				PoolSize:  10 * runtime.GOMAXPROCS(0),
			}),
		}))
		server.Use(requestid.New())
	}
	/*
		server.Post("/graphql", func(ctx *fiber.Ctx) error {
			body := new(GqlBody)

			if err := ctx.BodyParser(body); err != nil {
				return err
			}

			result := gql.Do(gql.Params{
				Context:        ctx.Context(),
				Schema:         TodoSchema,
				RequestString:  body.Query,
				VariableValues: body.Variables,
				OperationName:  body.Operation,
			})

			return ctx.JSON(result)
		})
	*/
	server.Static("/sandbox", "./public/sandbox.html")

	return Server{
		config: config,
		server: server,
	}
}

func (a *Server) start() error {
	if err := a.server.Listen(os.Getenv("SERVER_HOST") + ":" + os.Getenv("SERVER_PORT")); err != nil {
		log.Fatal(err)
	}

	return nil
}

func (s *Server) shutdown() error {
	if err := s.server.Shutdown(); err != nil {
		return err
	}

	return nil
}
