package mongodb

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/fx"
	"revil.dev-servers/lib"
)

func NewMongoDB(lc fx.Lifecycle, config *lib.Config) (*mongo.Database, error) {
	options := options.Client().ApplyURI(config.MongoDb.Uri)
	options.SetMaxPoolSize(100)
	options.SetMinPoolSize(10)
	options.SetMaxConnIdleTime(10 * time.Second)

	client, err := mongo.NewClient(options)
	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Print("Connect to MongoDB")
			return client.Connect(ctx)
		},
		OnStop: func(ctx context.Context) error {
			return client.Disconnect(ctx)
		},
	})

	return client.Database(config.MongoDb.Database), nil
}
