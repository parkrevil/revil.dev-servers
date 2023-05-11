package main

import (
	"context"
	"log"

	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/fx"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id"`
	Username  string             `bson:"username"`
	Password  string             `bson:"password"`
	ImageUrl  string             `bson:"image_url"`
	CreatedAt primitive.DateTime `bson:"created_at"`
}

var userType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"username": &graphql.Field{
			Type: graphql.String,
		},
		"imageUrl": &graphql.Field{
			Type: graphql.String,
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
			indexModel := mongo.IndexModel{
				Keys:    bson.D{{Key: "username", Value: -1}},
				Options: options.Index().SetUnique(true),
			}
			if _, err := collection.Indexes().CreateOne(ctx, indexModel); err != nil {
				return err
			}
			log.Print("Create MongoDB unique index: User.username")

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
				Type:        userType,
				Description: "내 정보",
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
	args := p.Args["input"]
	result, err := g.mongo.InsertOne(p.Context, User{
		Username: args["username"].(string),
		Password: args["password"].(string),
	})
	if err != nil {
		log.Print(err)
		return nil, err
	}

	log.Print(result)

	return nil, nil
}
