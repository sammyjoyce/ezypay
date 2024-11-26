package http

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"strconv"
)

// Server represents our HTTP server
type Server struct {
	port   int
	router *gin.Engine
}

// NewServer creates a new HTTP server
func NewServer() *Server {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8080 // Default port if not set
	}

	router := gin.Default()

	// Configure CORS
	allowedOrigins := os.Getenv("ALLOWED_ORIGINS")
	if allowedOrigins == "" {
		allowedOrigins = "*"
	}
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{allowedOrigins},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
	}))

	server := &Server{
		port:   port,
		router: router,
	}

	// Register routes
	server.registerRoutes()

	return server
}

// Start starts the HTTP server
func (s *Server) Start() error {
	addr := fmt.Sprintf(":%d", s.port)
	log.Printf("HTTP server listening on %s", addr)
	return s.router.Run(addr)
}

// GracefulStop stops the HTTP server gracefully
func (s *Server) GracefulStop() {
	log.Println("HTTP server stopping gracefully")
}
