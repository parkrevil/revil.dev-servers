package main

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

/*
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
*/
func main() {
	fx.New(
		fx.Provide(
			NewConfig,
			NewGraphQL,
			NewGgabongResolver,
			zap.NewProduction,
			NewHTTPServer,
		),
		fx.Invoke(func(*fiber.App) {}),
	).Run()
}
