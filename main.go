package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	gql "github.com/graphql-go/graphql"
	"github.com/joho/godotenv"
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
	env := os.Getenv("REVILDEV_ENV")
	if env == "" {
		log.Fatal("REVILDEV_ENV must be set")
	}

	envFilePath := ".env." + env
	if err := godotenv.Load(envFilePath); err != nil {
		log.Fatalf("Error loading %s file", envFilePath)
	}

	serverHost := os.Getenv("SERVER_HOST")
	serverPort := os.Getenv("SERVER_PORT")

	server := fiber.New(fiber.Config{
		AppName:       "revil.dev",
		Immutable:     true,
		CaseSensitive: true,
		StrictRouting: true,
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
	})

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

	go func() {
		if err := server.Listen(serverHost + ":" + serverPort); err != nil {
			log.Fatal(err)
		}
	}()

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
