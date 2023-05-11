package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/graphql-go/graphql"
)

type GqlBody struct {
	Query     string                 `json:"query"`
	Operation string                 `json:"operation"`
	Variables map[string]interface{} `json:"variables"`
}

type GraphQL struct {
	resolvers []GraphQLResolver
}

type GraphQLResolverSchema struct {
	Query    graphql.Fields
	Mutation graphql.Fields
}

type GraphQLResolver interface {
	GetSchemas() GraphQLResolverSchema
}

func NewGraphQL(ggabongResolver *GgabongResolver) *GraphQL {
	return &GraphQL{
		resolvers: []GraphQLResolver{
			ggabongResolver,
		},
	}
}

func (g *GraphQL) addHTTPHandler(server *fiber.App) error {
	/* 	queryFields := graphql.Fields{}
	   	mutationFields := graphql.Fields{}
	*/
	for _, resolver := range g.resolvers {
		schemas := resolver.GetSchemas()

		log.Print(schemas)
	}
	/*
		var query *graphql.Object
		var mutation *graphql.Object

		if len(queryFields) > 0 {
			query = graphql.NewObject(graphql.ObjectConfig{
				Name:   "Query",
				Fields: queryFields,
			})
		}

		if len(queryFields) > 0 {
			mutation = graphql.NewObject(graphql.ObjectConfig{
				Name:   "Mutation",
				Fields: mutationFields,
			})
		}

		schema, err := graphql.NewSchema(graphql.SchemaConfig{
			Query:    query,
			Mutation: mutation,
		})
		if err != nil {
			return err
		}

		server.Post("/graphql", func(ctx *fiber.Ctx) error {
			body := new(GqlBody)

			if err := ctx.BodyParser(body); err != nil {
				return err
			}

			result := graphql.Do(graphql.Params{
				Context:        ctx.Context(),
				Schema:         schema,
				RequestString:  body.Query,
				VariableValues: body.Variables,
				OperationName:  body.Operation,
			})

			return ctx.JSON(result)
		})
	*/
	server.Static("/sandbox", "./public/sandbox.html")

	return nil
}
