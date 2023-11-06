package trace

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"time"
)

func (c *Span) Info(msg interface{}, format ...string) {
	if len(format) > 0 {
		logrus.Infof(format[0], msg)
	} else {
		logrus.Info(msg)
	}
	c.span.AddEvent("@info",
		trace.WithTimestamp(time.Now()),
		trace.WithAttributes(
			attribute.String("msg", fmt.Sprintf("%v", msg)),
		))
}

func (c *Span) Warn(msg interface{}, format ...string) {
	if len(format) > 0 {
		logrus.Warnf(format[0], msg)
	} else {
		logrus.Warn(msg)
	}
	c.span.AddEvent("@warn",
		trace.WithTimestamp(time.Now()),
		trace.WithAttributes(
			attribute.String("msg", fmt.Sprintf("%v", msg)),
		),
	)
}

func (c *Span) Error(msg interface{}, format ...string) {
	if len(format) > 0 {
		logrus.Errorf(format[0], msg)
	} else {
		logrus.Error(msg)
	}
	c.span.AddEvent("@error",
		trace.WithTimestamp(time.Now()),
		trace.WithAttributes(
			attribute.String("msg", fmt.Sprintf("%v", msg)),
		),
	)
}

func (c *Span) Debug(msg interface{}, format ...string) {
	if len(format) > 0 {
		logrus.Debugf(format[0], msg)
	} else {
		logrus.Debug(msg)
	}
	c.span.AddEvent("@debug",
		trace.WithTimestamp(time.Now()),
		trace.WithAttributes(
			attribute.String("msg", fmt.Sprintf("%v", msg))),
	)
}
