package trace

import (
	"context"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/trace"
)

func (c *Trace) newJaegerExporter() (trace.SpanExporter, error) {
	if c.Insecure {
		return c.newInsecureJaegerExporter()
	}
	return c.newSecureJaegerExporter()
}

func (c *Trace) newInsecureJaegerExporter() (trace.SpanExporter, error) {
	if c.AccessToken == "" {
		return otlptracehttp.New(context.Background(),
			otlptracehttp.WithEndpoint(c.JaegerHost),
			otlptracehttp.WithInsecure(),
		)
	}

	return otlptracehttp.New(context.Background(),
		otlptracehttp.WithEndpoint(c.JaegerHost),
		otlptracehttp.WithInsecure(),
		otlptracehttp.WithHeaders(map[string]string{
			"Authorization": c.AccessToken,
		}),
	)
}

func (c *Trace) newSecureJaegerExporter() (trace.SpanExporter, error) {
	if c.AccessToken == "" {
		return otlptracehttp.New(context.Background(),
			otlptracehttp.WithEndpoint(c.JaegerHost),
		)
	}

	return otlptracehttp.New(context.Background(),
		otlptracehttp.WithEndpoint(c.JaegerHost),
		otlptracehttp.WithHeaders(map[string]string{
			"Authorization": c.AccessToken,
		}),
	)
}
