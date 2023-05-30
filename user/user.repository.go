package main

import (
	"context"

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

type userRepo struct {
	mongo *mongo.Collection
}

func newUserRepository(lc fx.Lifecycle, mongodb *mongo.Database) *userRepo {
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

	return &userRepo{
		mongo: collection,
	}
}
