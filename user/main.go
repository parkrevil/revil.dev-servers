package main

import (
	"go.uber.org/fx"
	"go.uber.org/zap"

	"revil.dev-servers/lib"
	"revil.dev-servers/lib/provider/mongodb"
	//	pb "revil.dev-servers/libs/services/article"
)

func main() {
	fx.New(
		fx.Provide(
			lib.NewConfig,
			mongodb.NewMongoDB,
			zap.NewProduction,
			NewServer,
			NewUserService,
		),
		fx.Invoke(func() {}),
	).Run()
}
