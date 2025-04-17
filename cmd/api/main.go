package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/adarsh-devappsys/rest-with-protobuf-logs/internal/app"
)

func main() {
	application, err := app.InitApp()
	if err != nil {
		log.Fatalf("Failed to initialize app: %v", err)
	}

	go func() {
		application.Run()
	}()
	// Set up a signal channel to handle graceful shutdown.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit // Wait for a termination signal.

	log.Println("Shutting down server...")
	log.Println("Application shut down gracefully.")
}
