package trace

import (
    "go.opentelemetry.io/otel/exporters/jaeger"
    "go.opentelemetry.io/otel/sdk/trace"
)

func newJaegerExporter(url string) (trace.SpanExporter, error) {
    return jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
}
