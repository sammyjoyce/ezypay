package handlers_test

import (
	"context"
	"net"
	"testing"

	"gitlab.com/australia-wide-first-aid/ezypay/internal/gen/hello"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func TestHelloHandler_SayHello(t *testing.T) {
	// Create a new handler
	handler := handlers.NewHelloHandler()

	// Test cases
	tests := []struct {
		name     string
		request  *hello.HelloRequest
		expected *hello.HelloResponse
		wantErr  bool
	}{
		{
			name: "basic greeting",
			request: &hello.HelloRequest{
				Name: "Alice",
			},
			expected: &hello.HelloResponse{
				Message: "Hello, Alice!, how are you",
			},
			wantErr: false,
		},
		{
			name: "empty name",
			request: &hello.HelloRequest{
				Name: "",
			},
			expected: &hello.HelloResponse{
				Message: "Hello, !, how are you",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Call the handler directly
			response, err := handler.SayHello(context.Background(), tt.request)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.expected.Message, response.Message)
		})
	}
}

func TestHelloHandler_Integration(t *testing.T) {
	// Create a buffered connection for testing
	listener := bufconn.Listen(1024 * 1024)

	// Create a new gRPC server
	server := grpc.NewServer()

	// Create and register the handler
	handler := handlers.NewHelloHandler()
	handler.RegisterWithServer(server)

	// Start the server
	go func() {
		if err := server.Serve(listener); err != nil {
			t.Errorf("Server exited with error: %v", err)
		}
	}()

	// Create a client connection
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return listener.Dial()
		}),
		grpc.WithInsecure(),
	)
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	// Create a client
	client := hello.NewHelloServiceClient(conn)

	// Test cases
	tests := []struct {
		name     string
		request  *hello.HelloRequest
		expected *hello.HelloResponse
		wantErr  bool
	}{
		{
			name: "integration test greeting",
			request: &hello.HelloRequest{
				Name: "Bob",
			},
			expected: &hello.HelloResponse{
				Message: "Hello, Bob!, how are you",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response, err := client.SayHello(ctx, tt.request)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.expected.Message, response.Message)
		})
	}

	// Clean up
	server.Stop()
}
