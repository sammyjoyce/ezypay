package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"{{go_module}}/pkg/grpc"
	"{{go_module}}/pkg/http"
)

func main() {
	// Create a context that will be canceled on SIGINT or SIGTERM
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Set up signal handling
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Create and start the gRPC server
	grpcServer := grpc.NewGRPCServer()
	go func() {
		if err := grpcServer.Start(); err != nil {
			log.Printf("gRPC server error: %v", err)
			cancel()
		}
	}()

	// Create and start the HTTP server
	httpServer := http.NewServer()
	go func() {
		if err := httpServer.Start(); err != nil {
			log.Printf("HTTP server error: %v", err)
			cancel()
		}
	}()

	// Wait for shutdown signal
	select {
	case <-sigChan:
		log.Println("Received shutdown signal")
	case <-ctx.Done():
		log.Println("Context canceled")
	}

	// Graceful shutdown
	httpServer.GracefulStop()
	grpcServer.GracefulStop()
}
