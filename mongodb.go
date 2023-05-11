package main

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/fx"
)

func NewMongoDB(lc fx.Lifecycle, config *Config) (*mongo.Database, error) {
	options := options.Client().ApplyURI(config.mongodb.uri)
	options.SetMaxPoolSize(100)
	options.SetMinPoolSize(10)
	options.SetMaxConnIdleTime(10 * time.Second)

	client, err := mongo.NewClient(options)
	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return client.Connect(ctx)
		},
		OnStop: func(ctx context.Context) error {
			return client.Disconnect(ctx)
		},
	})

	return client.Database(config.mongodb.database), nil
}
