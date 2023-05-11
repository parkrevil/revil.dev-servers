package main

import (
	"github.com/graphql-go/graphql"
)

var TodoList []Todo

type Todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

var todoType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Todo",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"text": &graphql.Field{
			Type: graphql.String,
		},
		"done": &graphql.Field{
			Type: graphql.Boolean,
		},
	},
})

type GgabongResolver struct{}

func NewGgabongResolver() *GgabongResolver {
	return &GgabongResolver{}
}

func (g *GgabongResolver) GetSchemas() GraphQLResolverSchema {
	return GraphQLResolverSchema{
		Query: graphql.Fields{
			"todo": &graphql.Field{
				Type:        todoType,
				Description: "Get single todo",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
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
			"todoList": &graphql.Field{
				Type:        graphql.NewList(todoType),
				Description: "List of todos",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return TodoList, nil
				},
			},
		},
	}
}
