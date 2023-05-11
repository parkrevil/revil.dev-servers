package main

import "github.com/graphql-go/graphql"

var createUserInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "CreateUserInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"username": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"password": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
})

type createUserInput struct {
	username string
	password string
}
