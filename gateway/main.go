package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	gql "github.com/graphql-go/graphql"
)

type GqlBody struct {
	Query     string                 `json:"query"`
	Operation string                 `json:"operation"`
	Variables map[string]interface{} `json:"variables"`
}

type Wiki struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
}

func main() {
	var wikis []Wiki

	content, err := ioutil.ReadFile("./data.json")

	if err != nil {
		log.Print("Error:", err)
	}

	err = json.Unmarshal(content, &wikis)

	if err != nil {
		log.Print("Error:", err)
	}

	wikiType := gql.NewObject(gql.ObjectConfig{
		Name: "Wiki",
		Fields: gql.Fields{
			"title": &gql.Field{
				Type: gql.String,
			},
			"description": &gql.Field{
				Type: gql.String,
			},
			"content": &gql.Field{
				Type: gql.String,
			},
		},
	})

	rootQuery := gql.NewObject(gql.ObjectConfig{
		Name: "RootQuery",
		Fields: gql.Fields{
			"wikiList": &gql.Field{
				Type:        gql.NewList(wikiType),
				Description: "List of wiki",
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					return wikis, nil
				},
			},
		},
	})
	rootMutation := gql.NewObject(gql.ObjectConfig{
		Name: "RootMutation",
		Fields: gql.Fields{
			"addWiki": &gql.Field{
				Type:        gql.Boolean, // the return type for this field
				Description: "add a new wiki",
				Args: gql.FieldConfigArgument{
					"title": &gql.ArgumentConfig{
						Type: gql.NewNonNull(gql.String),
					},
					"description": &gql.ArgumentConfig{
						Type: gql.NewNonNull(gql.String),
					},
					"content": &gql.ArgumentConfig{
						Type: gql.NewNonNull(gql.String),
					},
				},
				Resolve: func(params gql.ResolveParams) (interface{}, error) {
					title, _ := params.Args["name"].(string)
					description, _ := params.Args["description"].(string)
					content, _ := params.Args["imageUrl"].(string)

					wiki := Wiki{
						Title:       title,
						Description: description,
						Content:     content,
					}
					wikis = append(wikis, wiki)

					return true, nil
				},
			},
		},
	})
	schemaConfig := gql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	}
	schema, err := gql.NewSchema(schemaConfig)

	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	server := fiber.New(fiber.Config{
		AppName:           "revil.dev API Server",
		CaseSensitive:     true,
		EnablePrintRoutes: true,
		Immutable:         true,
		JSONEncoder:       json.Marshal,
		JSONDecoder:       json.Unmarshal,
	})

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

	if err := server.Listen(":20000"); err != nil {
		panic(err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	if err := server.ShutdownWithTimeout(2 * time.Second); err != nil {
		panic(err)
	}
}
