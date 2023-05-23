package main

import (
	"log"
	"net"
	"os"
	"os/signal"

	"revil.dev-servers/lib"
	//	pb "revil.dev-servers/libs/services/article"

	"google.golang.org/grpc"
)

func main() {
	fx.New(
		fx.Provide(
			lib.NewConfig,
			mongodb.NewMongoDB,
			zap.NewProduction,
		//	NewServer,
		//		user.NewUserService,
		),
		fx.Invoke(func() {}),
	).Run()
}
