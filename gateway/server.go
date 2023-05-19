package main

import (
	"context"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/goccy/go-json"
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/storage/redis/v2"
	"go.uber.org/fx"

	"revil.dev-servers/gateway/graph"
	"revil.dev-servers/gateway/graph/resolver"
	"revil.dev-servers/lib"
)

func NewHttpServer(lc fx.Lifecycle, config *lib.Config) (*fiber.App, error) {
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
	server.Use(limiter.New(limiter.Config{
		Max:               100,
		Expiration:        10 * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
		Storage: redis.New(redis.Config{
			Host:      config.Redis.Host,
			Port:      config.Redis.Port,
			Password:  config.Redis.Password,
			Database:  config.Redis.LimiterDb,
			Reset:     false,
			TLSConfig: nil,
			PoolSize:  10 * runtime.GOMAXPROCS(0),
		}),
	}))
	server.Use(requestid.New())

	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolver.Resolver{}}))

	server.Post("/gql", adaptor.HTTPHandlerFunc(h.ServeHTTP))
	server.Get("/", adaptor.HTTPHandlerFunc(playground.Handler("revil.dev GraphQL", "/gql")))

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go server.Listen(config.Server.Host + ":" + strconv.Itoa(config.Server.Port))

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.ShutdownWithContext(ctx)
		},
	})

	return server, nil
}
