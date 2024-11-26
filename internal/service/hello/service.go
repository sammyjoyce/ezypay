package hello

import (
	"context"
	"fmt"
)

// Service handles hello-related business logic
type Service struct{}

// NewService creates a new hello service
func NewService() *Service {
	return &Service{}
}

// SayHello handles the core greeting logic
func (s *Service) SayHello(ctx context.Context, name string) (string, error) {
	return fmt.Sprintf("Hello, %s!", name), nil
}
