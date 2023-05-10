package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	gql "github.com/graphql-go/graphql"
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
	/* 	ggabong := InitializeGgabong()
	   	ggabong.Start()

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
	*/

	app, err := InitializeApp()
	if err != nil {
		log.Fatal(err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	serverShutdown := make(chan struct{})

	go func() {
		<-quit

		log.Print("Shutting down...")
		log.Print("- fiber")

		if err := app.shutdown(); err != nil {
			log.Fatal(err)
		}

		serverShutdown <- struct{}{}
	}()

	app.start()

	<-serverShutdown

	fmt.Println("Shutted down")
}
