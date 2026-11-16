package app

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// GreetRequest is the JSON body accepted by POST /greet.
type GreetRequest struct {
	Name string `json:"name"`
}

// GreetResponse is the JSON body returned by POST /greet.
type GreetResponse struct {
	Message string `json:"message"`
}

// ErrorResponse is the JSON body returned on failure.
type ErrorResponse struct {
	Error string `json:"error"`
}

func greet(c *gin.Context) {
	var req GreetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid JSON body"})
		return
	}

	name := strings.TrimSpace(req.Name)
	if name == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "name is required"})
		return
	}

	c.JSON(http.StatusOK, GreetResponse{Message: "Hello, " + name + "!"})
}
