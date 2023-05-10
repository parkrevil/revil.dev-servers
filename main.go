package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"syscall"
	"time"
	"fmt"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/storage/redis/v2"
	gql "github.com/graphql-go/graphql"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	constant "revil.dev-servers/constant"
)

var TodoList []Todo

type Todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

var todoType = gql.NewObject(gql.ObjectConfig{
	Name: "Todo",
	Fields: gql.Fields{
		"id": &gql.Field{
			Type: gql.String,
		},
		"text": &gql.Field{
			Type: gql.String,
		},
		"done": &gql.Field{
			Type: gql.Boolean,
		},
	},
})

var rootQuery = gql.NewObject(gql.ObjectConfig{
	Name: "RootQuery",
	Fields: gql.Fields{
		"todo": &gql.Field{
			Type:        todoType,
			Description: "Get single todo",
			Args: gql.FieldConfigArgument{
				"id": &gql.ArgumentConfig{
					Type: gql.String,
				},
			},
			Resolve: func(params gql.ResolveParams) (interface{}, error) {
				idQuery, isOK := params.Args["id"].(string)
				if isOK {
					// Search for el with id
					for _, todo := range TodoList {
						if todo.ID == idQuery {
							return todo, nil
						}
					}
				}

				return Todo{}, nil
			},
		},
		"todoList": &gql.Field{
			Type:        gql.NewList(todoType),
			Description: "List of todos",
			Resolve: func(p gql.ResolveParams) (interface{}, error) {
				return TodoList, nil
			},
		},
	},
})

var TodoSchema, _ = gql.NewSchema(gql.SchemaConfig{
	Query: rootQuery,
})

type GqlBody struct {
	Query     string                 `json:"query"`
	Operation string                 `json:"operation"`
	Variables map[string]interface{} `json:"variables"`
}

func main() {
	e, err := InitializeEvent("test")
	if err != nil {
			fmt.Printf("failed to create event: %s\n", err)
			os.Exit(2)
	}
	e.Start()
	e.Start()

	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any

	url := "test"
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)

	env := os.Getenv("REVILDEV_ENV")
	if env == "" {
		log.Fatal("REVILDEV_ENV must be set")
	}

	envFilePath := ".env." + env
	if err := godotenv.Load(envFilePath); err != nil {
		log.Fatalf("Error loading %v file", envFilePath)
	}

	server := fiber.New(fiber.Config{
		AppName:       "revil.dev",
		Immutable:     true,
		CaseSensitive: true,
		StrictRouting: true,
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
	})

	/*
		TODO: Add custom logging middleware with Zap
	*/
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

	if env == constant.EnvProduction {
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

	if err := server.Listen(os.Getenv("SERVER_HOST") + ":" + os.Getenv("SERVER_PORT")); err != nil {
		log.Fatal(err)
	}

	server.Static("/sandbox", "./public/sandbox.html")

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Print("Shutting down...")
	log.Print("- fiber")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := server.ShutdownWithContext(ctx); err != nil {
		log.Fatal(err)
	}

	select {
	case <-ctx.Done():
		log.Print("Timeout shutting down fiber")
	}

	log.Print("done")
}
