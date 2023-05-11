package main

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.Provide(
			NewConfig,
			NewMongoDB,
			NewGraphQL,
			NewUserResolver,
			NewGgabongResolver,
			zap.NewProduction,
			NewHttpServer,
		),
		fx.Invoke(func(*fiber.App) {}),
	).Run()
}
