package main

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"revil.dev-servers/gateway/graph/resolver"
	//	"revil.dev-servers/gateway/user"
	"revil.dev-servers/lib"
	"revil.dev-servers/lib/provider/mongodb"
)

func main() {
	fx.New(
		fx.Provide(
			lib.NewConfig,
			mongodb.NewMongoDB,
			zap.NewProduction,
			NewHttpServer,
			resolver.NewResolver,
			//		user.NewUserService,
		),
		fx.Invoke(func(*fiber.App) {}),
	).Run()
}
