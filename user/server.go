package main

import (
	"net"

	"go.uber.org/zap"
)

func NewServer(lc fx.Lifecycle, config *lib.Config, log *zap.Logger) (*fiber.App, error) {
	sugar := logger.Sugar()
	l, err := net.Listen("tcp", config.UserServer.Port)

	if err != nil {
		sugar.Error("Failed to listen user server", err)
	}

	server := grpc.NewServer()

	pb.RegisterArticleServiceServer(server, NewArticleService())
	pb.RegisterArticlePageServiceServer(server, &ArticlePageService{})

	if err := server.Serve(l); err != nil {
		log.Fatalf("failed to serve slides server: %v", err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go server.Listen(config.Server.Host + ":" + strconv.Itoa(config.Server.Port))

			return nil
		},
		OnStop: func(ctx context.Context) error {
			server.GracefulStop()
			return server.ShutdownWithContext(ctx)
		},
	})

	return server, nil

}
