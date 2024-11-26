package http_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"{{go_module}}/pkg/http"
)

func TestHelloWorldHandler(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Create a new server instance
	server := NewServer()

	// Create a test HTTP recorder
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Call the handler
	server.HelloWorldHandler()(c)

	// Assert the response status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse the response body
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	// Assert the response content
	assert.Equal(t, "Hello, World!", response["message"])
}
