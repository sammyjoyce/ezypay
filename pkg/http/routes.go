package http

import (
	"net/http"

	"gitlab.com/australia-wide-first-aid/ezypay/internal/service/hello"

	"github.com/gin-gonic/gin"
)

// registerRoutes sets up all HTTP routes
func (s *Server) registerRoutes() {
	helloService := hello.NewService()

	s.router.GET("/", s.HelloWorldHandler())
	s.router.GET("/hello", s.SayHelloHandler(helloService))
	// Add more routes here as needed
}

// HelloWorldHandler handles the root endpoint
func (s *Server) HelloWorldHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	}
}

// SayHelloHandler handles hello requests
func (s *Server) SayHelloHandler(service *hello.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Query("name")
		if name == "" {
			name = "World"
		}

		message, err := service.SayHello(c.Request.Context(), name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": message,
		})
	}
}
