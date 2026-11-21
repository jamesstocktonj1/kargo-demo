package observability

import (
	"context"
	"os"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
)

func Init(r *gin.Engine) error {
	res, err := initResource(context.Background())
	if err != nil {
		return err
	}

	// Metrics
	if err := InitMetrics(res); err != nil {
		return err
	}

	mid, err := newObservabilityMiddleware()
	if err != nil {
		return err
	}
	r.Use(mid)

	return nil
}

func initResource(ctx context.Context) (*resource.Resource, error) {
	return resource.New(ctx,
		resource.WithAttributes(
			attribute.String("service.name", "kargo-demo"),
			attribute.String("environment", os.Getenv("ENV")),
		),
		resource.WithHost(),
		resource.WithFromEnv(),
	)
}
