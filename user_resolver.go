package main

import (
	"context"
	"time"

	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/fx"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Username  string             `bson:"username"`
	Password  string             `bson:"password"`
	Nickname  string             `bson:"nickname"`
	Email     string             `bson:"email"`
	ImageUrl  string             `bson:"image_url"`
	CreatedAt primitive.DateTime `bson:"created_at,omitempty"`
	UpdatedAt primitive.DateTime `bson:"updated_at,omitempty"`
}

var userType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
		},
		"username": &graphql.Field{
			Type: graphql.String,
		},
		"nickname": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"imageUrl": &graphql.Field{
			Type: graphql.String,
		},
		"createdAt": &graphql.Field{
			Type: graphql.DateTime,
		},
	},
})

type UserResolver struct {
	mongo *mongo.Collection
}

func NewUserResolver(lc fx.Lifecycle, mongodb *mongo.Database) *UserResolver {
	collection := mongodb.Collection("user")

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			indexModels := []mongo.IndexModel{
				{
					Keys:    bson.D{{Key: "username", Value: -1}},
					Options: options.Index().SetUnique(true),
				},
				{
					Keys:    bson.D{{Key: "nickname", Value: -1}},
					Options: options.Index().SetUnique(true),
				},
				{
					Keys:    bson.D{{Key: "email", Value: -1}},
					Options: options.Index().SetUnique(true),
				},
			}
			if _, err := collection.Indexes().CreateMany(ctx, indexModels); err != nil {
				return err
			}

			return nil
		},
	})

	return &UserResolver{
		mongo: collection,
	}
}

func (g *UserResolver) GetSchemas() GraphQLResolverSchema {
	return GraphQLResolverSchema{
		Mutation: graphql.Fields{
			"createUser": &graphql.Field{
				Type:        graphql.Boolean,
				Description: "사용자 생성",
				Args: graphql.FieldConfigArgument{
					"input": &graphql.ArgumentConfig{
						Type: createUserInputType,
					},
				},
				Resolve: g.create,
			},
		},
	}
}

func (g *UserResolver) create(p graphql.ResolveParams) (interface{}, error) {
	input := p.Args["input"].(map[string]interface{})

	_, err := g.mongo.InsertOne(p.Context, User{
		Username:  input["username"].(string),
		Password:  input["password"].(string),
		Nickname:  input["nickname"].(string),
		Email:     input["email"].(string),
		ImageUrl:  input["imageUrl"].(string),
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
	})
	if err != nil {
		return false, err
	}

	return true, nil
}
