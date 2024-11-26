package handlers

import (
	"context"
	"log"

	"google.golang.org/grpc"
	pb "gitlab.com/australia-wide-first-aid/ezypay/internal/gen/hello"
	helloservice "gitlab.com/australia-wide-first-aid/ezypay/internal/service/hello"
)

// HelloHandler handles gRPC requests for the Hello service
type HelloHandler struct {
	pb.UnimplementedHelloServiceServer
	service *helloservice.Service
}

// NewHelloHandler creates a new HelloHandler
func NewHelloHandler() *HelloHandler {
	return &HelloHandler{
		service: helloservice.NewService(),
	}
}

// RegisterWithServer implements the Handler interface
func (h *HelloHandler) RegisterWithServer(server interface{}) {
	if grpcServer, ok := server.(*grpc.Server); ok {
		pb.RegisterHelloServiceServer(grpcServer, h)
	}
}

// SayHello implements the HelloService gRPC method
func (h *HelloHandler) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received: %s", req.GetName())

	message, err := h.service.SayHello(ctx, req.GetName())
	if err != nil {
		return nil, err
	}

	return &pb.HelloResponse{
		Message: message,
	}, nil
}
