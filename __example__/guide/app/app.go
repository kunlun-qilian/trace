package app

import (
    "context"
    "fmt"
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/attribute"
    "go.opentelemetry.io/otel/trace"
    "io"
    "log"
    "strconv"
    "time"
)

const name = "fib"

// App is a Fibonacci computation application.
type App struct {
    r io.Reader
    l *log.Logger
}

// NewApp returns a new App.
func NewApp(r io.Reader, l *log.Logger) *App {
    return &App{r: r, l: l}
}

// Run starts polling users for Fibonacci number requests and writes results.
func (a *App) Run(ctx context.Context) error {
    for {
        // Each execution of the run loop, we should get a new "root" span and context.
        newCtx, span := otel.Tracer(name).Start(ctx, "Run", trace.WithTimestamp(time.Now()))

        n, err := a.Poll(newCtx)
        if err != nil {
            span.End(trace.WithTimestamp(time.Now()))
            return err
        }

        a.Write(newCtx, n)
        span.End(trace.WithTimestamp(time.Now()))
    }
}

// Poll asks a user for input and returns the request.
func (a *App) Poll(ctx context.Context) (uint, error) {
    _, span := otel.Tracer(name).Start(ctx, "Poll", trace.WithTimestamp(time.Now()))
    defer span.End(trace.WithTimestamp(time.Now()))

    a.l.Print("What Fibonacci number would you like to know: ")

    var n uint
    _, err := fmt.Fscanf(a.r, "%d\n", &n)

    // Store n as a string to not overflow an int64.
    nStr := strconv.FormatUint(uint64(n), 10)
    span.SetAttributes(attribute.String("request.n", nStr))

    return n, err
}

// Write writes the n-th Fibonacci number back to the user.
func (a *App) Write(ctx context.Context, n uint) {
    var span trace.Span
    ctx, span = otel.Tracer(name).Start(ctx, "Write", trace.WithTimestamp(time.Now()))
    defer span.End(trace.WithTimestamp(time.Now()))

    f, err := func(ctx context.Context) (uint64, error) {
        _, span := otel.Tracer(name).Start(ctx, "Fibonacci")
        defer span.End(trace.WithTimestamp(time.Now()))
        return Fibonacci(n)
    }(ctx)
    if err != nil {
        a.l.Printf("Fibonacci(%d): %v\n", n, err)
    } else {
        a.l.Printf("Fibonacci(%d) = %d\n", n, f)
    }
}
