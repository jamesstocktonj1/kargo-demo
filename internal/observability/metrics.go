package observability

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	otelprom "go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
)

func InitMetrics(r *resource.Resource) error {
	reg := prometheus.NewRegistry()
	exp, err := otelprom.New(
		otelprom.WithRegisterer(reg),
		otelprom.WithoutTargetInfo(),
		otelprom.WithoutScopeInfo(),
		otelprom.WithResourceAsConstantLabels(attribute.NewAllowKeysFilter(
			"environment",
			"service.version",
			"host.name",
		)),
	)
	if err != nil {
		return err
	}

	provider := metric.NewMeterProvider(
		metric.WithReader(exp),
		metric.WithResource(r),
	)
	otel.SetMeterProvider(provider)

	go func() {
		metrics := http.NewServeMux()
		metrics.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
		http.ListenAndServe(":9090", metrics)
	}()

	return nil
}
