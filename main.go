package main

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"revil.dev-servers/graph/resolver"
)

func main() {
	fx.New(
		fx.Provide(
			NewConfig,
			NewMongoDB,
			NewUserResolver,
			zap.NewProduction,
			NewHttpServer,
			resolver.NewResolver,
		),
		fx.Invoke(func(*fiber.App) {}),
	).Run()
}
