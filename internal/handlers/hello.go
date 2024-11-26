package handlers

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"{{go_module}}/internal/gen/hello"
	"{{go_module}}/internal/service/hello"
)

// HelloHandler handles gRPC requests for the Hello service
type HelloHandler struct {
	hello.UnimplementedHelloServiceServer
	service *hello.Service
}

// NewHelloHandler creates a new HelloHandler
func NewHelloHandler() *HelloHandler {
	return &HelloHandler{
		service: hello.NewService(),
	}
}

// RegisterWithServer implements the Handler interface
func (h *HelloHandler) RegisterWithServer(server interface{}) {
	if grpcServer, ok := server.(*grpc.Server); ok {
		hello.RegisterHelloServiceServer(grpcServer, h)
	}
}

// SayHello implements the HelloService gRPC method
func (h *HelloHandler) SayHello(ctx context.Context, req *hello.HelloRequest) (*hello.HelloResponse, error) {
	log.Printf("Received: %s", req.GetName())

	message, err := h.service.SayHello(ctx, req.GetName())
	if err != nil {
		return nil, err
	}

	return &hello.HelloResponse{
		Message: message,
	}, nil
}
