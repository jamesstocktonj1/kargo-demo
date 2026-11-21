package observability

import (
	"os"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
)

func Init(r *gin.Engine) error {
	res := initResource()

	// Metrics
	err := InitMetrics(res)
	if err != nil {
		return err
	}

	mid, err := newObservabilityMiddleware()
	if err != nil {
		return err
	}
	r.Use(mid)

	return nil
}

func initResource() *resource.Resource {
	return resource.NewSchemaless(
		attribute.String("environment", os.Getenv("ENV")),
	)
}
