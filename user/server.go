package main

import (
	"context"
	"net"
	"strconv"

	"go.uber.org/fx"
	"google.golang.org/grpc"
	"revil.dev-servers/lib"
	pb "revil.dev-servers/lib/service/user"

	"go.uber.org/zap"
)

func NewServer(lc fx.Lifecycle, config *lib.Config, logger *zap.Logger, userService *UserService) (*grpc.Server, error) {
	listenAddress := config.UserGrpcServer.Host + ":" + strconv.Itoa(config.UserGrpcServer.Port)
	l, err := net.Listen("tcp", listenAddress)
	if err != nil {
		logger.Error("Failed to listen user server", zap.Error(err))
	}

	server := grpc.NewServer()

	pb.RegisterUserServiceServer(server, userService)

	if err := server.Serve(l); err != nil {
		logger.Error("failed to serve slides server: %v", zap.Error(err))
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if err := server.Serve(l); err != nil {
				logger.Error("failed to serve user server", zap.Error(err))
			}

			return nil
		},
		OnStop: func(ctx context.Context) error {
			server.GracefulStop()

			return nil
		},
	})

	return server, nil
}
