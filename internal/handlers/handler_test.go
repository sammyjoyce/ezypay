package handlers_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"gitlab.com/australia-wide-first-aid/ezypay/internal/handlers"
)

// mockHandler implements the Handler interface for testing
type mockHandler struct {
	registerCalled bool
	server         interface{}
}

func (m *mockHandler) RegisterWithServer(server interface{}) {
	m.registerCalled = true
	m.server = server
}

func TestRegistry(t *testing.T) {
	t.Run("NewRegistry creates empty registry", func(t *testing.T) {
		registry := handlers.NewRegistry()
		assert.NotNil(t, registry)
		assert.Empty(t, registry.handlers)
	})

	t.Run("Register adds handler to registry", func(t *testing.T) {
		registry := handlers.NewRegistry()
		handler := &mockHandler{}

		registry.Register(handler)

		assert.Len(t, registry.handlers, 1)
		assert.Contains(t, registry.handlers, handler)
	})

	t.Run("RegisterAllWithServer calls RegisterWithServer on all handlers", func(t *testing.T) {
		registry := handlers.NewRegistry()
		handler1 := &mockHandler{}
		handler2 := &mockHandler{}

		registry.Register(handler1)
		registry.Register(handler2)

		server := grpc.NewServer()
		registry.RegisterAllWithServer(server)

		assert.True(t, handler1.registerCalled)
		assert.True(t, handler2.registerCalled)
		assert.Equal(t, server, handler1.server)
		assert.Equal(t, server, handler2.server)
	})
}
