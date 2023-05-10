package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	env := os.Getenv("REVILDEV_ENV")
	if env == "" {
		log.Fatal("REVILDEV_ENV must be set")
	}

	envFilePath := ".env." + env
	if err := godotenv.Load(envFilePath); err != nil {
		log.Fatalf("Error loading %s file", envFilePath)
	}

	serverHost := os.Getenv("SERVER_HOST")
	serverPort := os.Getenv("SERVER_PORT")

	server := fiber.New(fiber.Config{
		AppName:       "revil.dev",
		Immutable:     true,
		CaseSensitive: true,
		StrictRouting: true,
	})

	go func() {
		if err := server.Listen(serverHost + ":" + serverPort); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Print("Shutting down...")
	log.Print("- fiber")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := server.ShutdownWithContext(ctx); err != nil {
		log.Fatal(err)
	}

	select {
	case <-ctx.Done():
		log.Print("Timeout shutting down fiber")
	}

	log.Print("done")
}
