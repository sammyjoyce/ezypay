package grpc

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"gitlab.com/australia-wide-first-aid/ezypay/internal/handlers"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// GRPCServer represents our gRPC server
type GRPCServer struct {
	port     int
	host     string
	server   *grpc.Server
	registry *handlers.Registry
}

// NewGRPCServer creates a new GRPCServer
func NewGRPCServer() *GRPCServer {
	port, err := strconv.Atoi(os.Getenv("GRPC_PORT"))
	if err != nil {
		port = 50051 // Default port if not set
	}
	host := os.Getenv("GRPC_HOST")
	if host == "" {
		host = "0.0.0.0" // Default to all interfaces if not set
	}

	srv := grpc.NewServer()
	registry := handlers.NewRegistry()

	// Register all handlers
	registry.Register(handlers.NewHelloHandler())

	grpcServer := &GRPCServer{
		port:     port,
		host:     host,
		server:   srv,
		registry: registry,
	}

	// Register all handlers with the gRPC server
	registry.RegisterAllWithServer(srv)
	reflection.Register(srv)

	return grpcServer
}

// Start starts the gRPC server
func (s *GRPCServer) Start() error {
	addr := fmt.Sprintf("%s:%d", s.host, s.port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	log.Printf("gRPC server listening on %s", addr)
	return s.server.Serve(lis)
}

// GracefulStop stops the gRPC server gracefully
func (s *GRPCServer) GracefulStop() {
	log.Println("gRPC server stopping gracefully")
	s.server.GracefulStop()
}
