package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/golang/protobuf/ptypes/empty"
	gql "github.com/graphql-go/graphql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "revil.dev-servers/libs/services/article"
)

type GqlBody struct {
	Query     string                 `json:"query"`
	Operation string                 `json:"operation"`
	Variables map[string]interface{} `json:"variables"`
}

func main() {
	articleServiceConn, err := grpc.Dial("localhost:20001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to init article service connection: %v", err)
	}
	defer articleServiceConn.Close()
	articleService := pb.NewArticleServiceClient(articleServiceConn)

	articleType := gql.NewObject(gql.ObjectConfig{
		Name: "Article",
		Fields: gql.Fields{
			"id": &gql.Field{
				Type: gql.String,
			},
			"title": &gql.Field{
				Type: gql.String,
			},
			"description": &gql.Field{
				Type: gql.String,
			},
		},
	})

	getArticlesQuery := gql.NewObject(gql.ObjectConfig{
		Name: "GetArticles",
		Fields: gql.Fields{
			"articles": &gql.Field{
				Type:        gql.NewList(articleType),
				Description: "Get articles",
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					ctx, cancel := context.WithTimeout(context.Background(), time.Second)
					defer cancel()

					articles, err := articleService.GetArticles(ctx, &empty.Empty{})
					if err != nil {
						log.Fatalf("GetArticles failed: %v", err)
					}

					return articles.Articles, nil
				},
			},
		},
	})
	/* 	rootMutation := gql.NewObject(gql.ObjectConfig{
		Name: "RootMutation",
		Fields: gql.Fields{
			"createArticle": &gql.Field{
				Type:        gql.Boolean,
				Description: "Create a new article",
				Args: gql.FieldConfigArgument{
					"title": &gql.ArgumentConfig{
						Type: gql.NewNonNull(gql.String),
					},
					"description": &gql.ArgumentConfig{
						Type: gql.NewNonNull(gql.String),
					},
				},
				Resolve: func(params gql.ResolveParams) (interface{}, error) {
					id, _ := params.Args["id"].(string)
					title, _ := params.Args["title"].(string)
					description, _ := params.Args["description"].(string)

					article := Article{
						Id:          id,
						Title:       title,
						Description: description,
					}
					articles = append(articles, article)

					return true, nil
				},
			},
		},
	}) */
	schemaConfig := gql.SchemaConfig{
		Query: getArticlesQuery,
		//Mutation: rootMutation,
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
