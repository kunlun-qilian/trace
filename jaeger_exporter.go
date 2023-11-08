package trace

import (
	"context"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/trace"
)

func (c *Trace) newJaegerExporter() (trace.SpanExporter, error) {
	if c.Insecure {
		return otlptracehttp.New(context.Background(),
			otlptracehttp.WithEndpoint(c.JaegerHost),
			otlptracehttp.WithInsecure(),
		)
	}
	return otlptracehttp.New(context.Background(),
		otlptracehttp.WithEndpoint(c.JaegerHost),
	)
}
