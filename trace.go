package trace

import (
	"context"
	b3prop "go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/otel"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	oteltrace "go.opentelemetry.io/otel/trace"
	"reflect"
	"time"
)

var ServiceName string

type Trace struct {
	// jaeger 地址 本地调试时或使用外部jaeger的时候,需要配置对应的 jaegerHost,留空默认访问本集群的jaeger
	JaegerHost string `env:""`
	// AlwaysSample default false
	AlwaysSample bool `env:""`
	// 默认启用证书，关闭证书设置为true
	Insecure bool `env:""`
	// AccessToken  访问Jaeger的access token,可以为空
	AccessToken string `env:""`
	ServiceName string
}

func (c *Trace) SetDefaults() {
	if c.JaegerHost == "" {
		c.JaegerHost = "jaeger-otlp.observability:4318"
	}
}

func (c *Trace) Init() {
	c.SetDefaults()

	ServiceName = c.ServiceName

	exporter, err := c.newJaegerExporter()
	if err != nil {
		panic(err)
	}

	var opts []sdktrace.TracerProviderOption
	if c.AlwaysSample {
		opts = append(opts, sdktrace.WithBatcher(exporter), sdktrace.WithSampler(sdktrace.AlwaysSample()),
			sdktrace.WithResource(newResource(ServiceName)))
	} else {
		opts = append(opts, sdktrace.WithBatcher(exporter),
			sdktrace.WithResource(newResource(ServiceName)))
	}

	tp := sdktrace.NewTracerProvider(opts...)
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

func (c *Span) TraceSpan() oteltrace.Span {
	return c.span
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
