module github.com/kunlun-qilian/trace

go 1.20

require (
	github.com/sirupsen/logrus v1.9.3
	go.opentelemetry.io/contrib/propagators/b3 v1.20.0
	go.opentelemetry.io/otel v1.19.0
	go.opentelemetry.io/otel/exporters/jaeger v1.17.0
	go.opentelemetry.io/otel/sdk v1.19.0
	go.opentelemetry.io/otel/trace v1.19.0
)

require (
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	go.opentelemetry.io/otel/metric v1.19.0 // indirect
	golang.org/x/sys v0.12.0 // indirect
)
