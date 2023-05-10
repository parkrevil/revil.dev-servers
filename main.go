package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	server := fiber.New(fiber.Config{
		AppName:       "revil.dev",
		Immutable:     true,
		CaseSensitive: true,
		StrictRouting: true,
	})

	go func() {
		if err := server.Listen(":3000"); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Print("Shutting down...")
	log.Print("- fiber")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
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
