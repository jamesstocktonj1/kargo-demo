package app

import (
	"github.com/gin-gonic/gin"

	"github.com/jamesstocktonj1/kargo-demo/internal/observability"
)

// NewRouter builds the gin engine with all routes registered. It panics if
// metrics cannot be initialised, since the server should not start without them.
func NewRouter() (*gin.Engine, error) {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	if err := observability.Init(r); err != nil {
		return nil, err
	}

	r.POST("/greet", greet)

	return r, nil
}
