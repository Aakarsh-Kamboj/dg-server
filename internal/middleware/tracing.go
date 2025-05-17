package middleware

import (
	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.uber.org/zap"
)

// InitTracer initializes OpenTelemetry Tracing
func InitTracer(logger *zap.Logger) (*trace.TracerProvider, error) {
	exporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		logger.Fatal("Failed to create stdout exporter", zap.Error(err))
	}

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resource.NewWithAttributes(semconv.SchemaURL, semconv.ServiceNameKey.String("api-gateway"))),
	)
	otel.SetTracerProvider(tp)

	logger.Info("Tracing initialized")

	return tp, nil
}

// InitMetrics initializes OpenTelemetry Metrics (Prometheus)
func InitMetrics(logger *zap.Logger) (*metric.MeterProvider, error) {
	exporter, err := prometheus.New()
	if err != nil {
		logger.Fatal("Failed to create Prometheus exporter", zap.Error(err))
	}

	mp := metric.NewMeterProvider(metric.WithReader(exporter))
	otel.SetMeterProvider(mp)

	logger.Info("Metrics initialized")

	return mp, nil
}

// TracingMiddleware applies OpenTelemetry Tracing to Echo
func TracingMiddleware(e *echo.Echo) {
	e.Use(otelecho.Middleware("api-gateway"))
}
