package grpc

import (
	"context"
	"net"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"{{go_module}}/internal/gen/hello"
)

func TestGRPCServer(t *testing.T) {
	t.Run("NewGRPCServer initializes with default values", func(t *testing.T) {
		// Clear environment variables
		os.Unsetenv("GRPC_PORT")
		os.Unsetenv("GRPC_HOST")

		server := NewGRPCServer()
		assert.NotNil(t, server)
		assert.Equal(t, 50051, server.port)
		assert.Equal(t, "0.0.0.0", server.host)
		assert.NotNil(t, server.server)
		assert.NotNil(t, server.registry)
	})

	t.Run("NewGRPCServer uses environment variables", func(t *testing.T) {
		// Set environment variables
		os.Setenv("GRPC_PORT", "50052")
		os.Setenv("GRPC_HOST", "localhost")
		defer func() {
			os.Unsetenv("GRPC_PORT")
			os.Unsetenv("GRPC_HOST")
		}()

		server := NewGRPCServer()
		assert.Equal(t, 50052, server.port)
		assert.Equal(t, "localhost", server.host)
	})

	t.Run("Server starts and handles requests", func(t *testing.T) {
		// Create a buffered connection for testing
		listener := bufconn.Listen(1024 * 1024)

		// Create and start the server
		server := NewGRPCServer()
		go func() {
			if err := server.server.Serve(listener); err != nil {
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

		// Test the hello service
		resp, err := client.SayHello(ctx, &hello.HelloRequest{Name: "Test"})
		assert.NoError(t, err)
		assert.Equal(t, "Hello, Test!, how are you", resp.Message)

		// Clean up
		server.GracefulStop()
	})

	t.Run("Server starts and stops gracefully", func(t *testing.T) {
		server := NewGRPCServer()

		// Start the server in a goroutine
		go func() {
			if err := server.Start(); err != nil {
				// Ignore address already in use errors during testing
				if opErr, ok := err.(*net.OpError); !ok || opErr.Op != "listen" {
					t.Errorf("Server exited with unexpected error: %v", err)
				}
			}
		}()

		// Give the server time to start
		time.Sleep(100 * time.Millisecond)

		// Stop the server gracefully
		server.GracefulStop()
	})
}
