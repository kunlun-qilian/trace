package app

import (
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/trace"
)

func NewJaegerExporter(url string) (trace.SpanExporter, error) {
	urlSuffix := "/api/traces"
	return jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url + urlSuffix)))
}
