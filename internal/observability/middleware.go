package observability

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

var (
	meter = otel.Meter("kargo-demo")
)

func newObservabilityMiddleware() (gin.HandlerFunc, error) {
	count, err := meter.Int64Counter("http_response_status_code",
		metric.WithDescription("Count of HTTP responses by path, method and status code"),
	)
	if err != nil {
		return nil, err
	}

	duration, err := meter.Float64Histogram("http_server_request_duration_milliseconds",
		metric.WithDescription("HTTP server request duration in milliseconds"),
	)
	if err != nil {
		return nil, err
	}

	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		elapsed := time.Since(start)

		opts := metric.WithAttributes(
			attribute.String("path", c.Request.URL.Path),
			attribute.String("method", c.Request.Method),
			attribute.Int("status_code", c.Writer.Status()),
		)

		count.Add(c, 1, opts)
		duration.Record(c, float64(elapsed.Microseconds())/1000, opts)
	}, nil
}
