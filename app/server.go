package app

import "github.com/gin-gonic/gin"

// NewRouter builds the gin engine with all routes registered.
func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	r.POST("/greet", greet)

	return r
}
