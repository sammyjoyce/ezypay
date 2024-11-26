package handlers

// Handler is the interface that all gRPC handlers must implement
type Handler interface {
	// RegisterWithServer registers the handler with the appropriate gRPC server
	RegisterWithServer(server interface{})
}

// Registry maintains a collection of all gRPC handlers
type Registry struct {
	handlers []Handler
}

// NewRegistry creates a new handler registry
func NewRegistry() *Registry {
	return &Registry{
		handlers: make([]Handler, 0),
	}
}

// Register adds a new handler to the registry
func (r *Registry) Register(handler Handler) {
	r.handlers = append(r.handlers, handler)
}

// RegisterAllWithServer registers all handlers with the gRPC server
func (r *Registry) RegisterAllWithServer(server interface{}) {
	for _, handler := range r.handlers {
		handler.RegisterWithServer(server)
	}
}
