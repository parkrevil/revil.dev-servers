package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

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
	app := NewApp()
	defer app.Shutdown()

	if err := app.Start(); err != nil {
		panic(err)
	}

/* 	articleServiceConn, err := grpc.Dial("localhost:20001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to init article service connection: %v", err)
	}
	defer articleServiceConn.Close()
	articleService := pb.NewArticleServiceClient(articleServiceConn)
 */
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

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	if err := app.Shutdown(); err != nil {
		panic(err)
	}
}
