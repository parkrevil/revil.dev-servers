package main

import "github.com/graphql-go/graphql"

var createUserInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name:        "CreateUserInput",
	Description: "사용자 생성 입력 정보",
	Fields: graphql.InputObjectConfigFieldMap{
		"username": &graphql.InputObjectFieldConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "사용자 아이디",
		},
		"password": &graphql.InputObjectFieldConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "사용자 비밀번호",
		},
		"nickname": &graphql.InputObjectFieldConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "사용자 닉네임",
		},
		"email": &graphql.InputObjectFieldConfig{
			Type:        graphql.NewNonNull(graphql.NewScalar()),
			Description: "사용자 이메일",
		},
		"imageUrl": &graphql.InputObjectFieldConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "사용자 이미지 URL",
		},
	},
})
