package trace

import (
	"context"
	"fmt"
	b3prop "go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/otel"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	oteltrace "go.opentelemetry.io/otel/trace"
	"reflect"
	"time"
)

var ServiceName string

type Trace struct {
	// jaeger 地址
	JaegerHost string
	// 默认为false,本地调试时,需要配置对应的 jaeger-collector
	Local bool `env:""`
	// AlwaysSample default false
	AlwaysSample bool `env:""`

	ServiceName string

	jaegerUrl string
}

func (c *Trace) SetDefaults() {
	if !c.Local {
		c.JaegerHost = "http://jaeger-collector.observability:14268"
	}
	c.jaegerUrl = fmt.Sprintf("%s/api/traces", c.JaegerHost)
}

func (c *Trace) Init() {
	c.SetDefaults()

	ServiceName = c.ServiceName

	exporter, err := newJaegerExporter(c.jaegerUrl)
	if err != nil {
		panic(err)
	}

	tp := &sdktrace.TracerProvider{}
	if c.AlwaysSample {
		tp = sdktrace.NewTracerProvider(
			sdktrace.WithBatcher(exporter),
			sdktrace.WithSampler(sdktrace.AlwaysSample()),
			sdktrace.WithResource(newResource(c.ServiceName)),
		)
	} else {
		tp = sdktrace.NewTracerProvider(
			sdktrace.WithBatcher(exporter),
			sdktrace.WithResource(newResource(c.ServiceName)),
		)
	}

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(b3prop.New())
}

func NewSpan(ctx context.Context, span oteltrace.Span) *Span {
	return &Span{
		ctx:  ctx,
		span: span,
	}
}

type Span struct {
	ctx  context.Context
	span oteltrace.Span
}

func (c *Span) Context() context.Context {
	return c.ctx
}

func (c *Span) End() {
	defer c.span.End(oteltrace.WithTimestamp(time.Now()))
}

func Start(ctx context.Context, spanName string, opts ...oteltrace.SpanStartOption) *Span {
	traceCtx, span := otel.Tracer(ServiceName).Start(ctx, spanName, opts...)
	return &Span{
		ctx:  traceCtx,
		span: span,
	}
}

type ContextTraceSpan struct {
}

var ContextTraceSpanKey = reflect.TypeOf(ContextTraceSpan{}).String()

func GetTraceSpanFromContext(ctx context.Context) *Span {
	return ctx.Value(ContextTraceSpanKey).(*Span)
}
