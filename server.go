package main

import (
	"context"
	"strconv"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

func NewHTTPServer(lc fx.Lifecycle, config *Config) *fiber.App {
	server := fiber.New(fiber.Config{
		AppName:       "revil.dev",
		Immutable:     true,
		CaseSensitive: true,
		StrictRouting: true,
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
	})

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go server.Listen(config.server.host + ":" + strconv.Itoa(config.server.port))

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.ShutdownWithContext(ctx)
		},
	})

	return server
}

/*
func NewHTTPServer(lc fx.Lifecycle) *Server {
	app := fiber.New(fiber.Config{
		AppName:       "revil.dev",
		Immutable:     true,
		CaseSensitive: true,
		StrictRouting: true,
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
	})
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     os.Getenv("SERVER_CORS_ORIGINS"),
		AllowHeaders:     "Referer, Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "POST,GET",
		AllowCredentials: false,
	}))
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestCompression,
	}))
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))

	if config.env == Production {
		redisPort, err := strconv.Atoi(os.Getenv("REDIS_PORT"))
		if err != nil {
			log.Fatalf("Invalid redis port: %v", err)
		}

		app.Use(limiter.New(limiter.Config{
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
		app.Use(requestid.New())
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
/* 	app.Static("/sandbox", "./public/sandbox.html")

	server := Server{
		config: config,
		server: app,
	}

	lc.Append(fx.Hook{
    OnStart: func(ctx context.Context) error {
			if err := server.server.Listen(os.Getenv("SERVER_HOST") + ":" + os.Getenv("SERVER_PORT")); err != nil {
				log.Fatal(err)
			}

			return nil

      ln, err := net.Listen("tcp", srv.Addr)
      if err != nil {
        return err
      }
      fmt.Println("Starting HTTP server at", srv.Addr)
      go srv.Serve(ln)
      return nil
    },
    OnStop: func(ctx context.Context) error {
      return srv.Shutdown(ctx)
    },
  })

	return &
}

func (s *Server) shutdown() error {
	if err := s.server.Shutdown(); err != nil {
		return err
	}

	return nil
}
*/
