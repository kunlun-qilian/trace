package app

import (
	"context"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/trace"
)

func NewJaegerExporter(url string) (trace.SpanExporter, error) {
	return otlptracehttp.New(context.Background(),
		otlptracehttp.WithEndpoint(url),
	)
}
