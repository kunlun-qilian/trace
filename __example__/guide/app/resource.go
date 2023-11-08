package app

import (
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
)

// newResource returns a resource describing this application.
func NewResource(serviceName string) *resource.Resource {
	return resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(serviceName),
		semconv.TelemetrySDKLanguageGo,
		semconv.TelemetrySDKNameKey.String("opentelemetry"),
	)
}
