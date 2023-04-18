package main

import (
	"log"
	"net"
	"os"
	"os/signal"

	pb "revil.dev-servers/libs/services/article"

	"google.golang.org/grpc"
)

func main() {
	l, err := net.Listen("tcp", ":20001")

	if err != nil {
		log.Fatalf("failed to listen slides serdver: %v", err)
	}

	server := grpc.NewServer()

	pb.RegisterArticleServiceServer(server, &ArticleService{})
	pb.RegisterArticlePageServiceServer(server, &ArticlePageService{})

	if err := server.Serve(l); err != nil {
		log.Fatalf("failed to serve slides server: %v", err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	server.GracefulStop()
}
