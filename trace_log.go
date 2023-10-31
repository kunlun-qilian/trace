package trace

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

func Info(span *Span, msg interface{}, format ...string) {
	if len(format) > 0 {
		span.span.RecordError(fmt.Errorf(format[0], msg))
		logrus.Infof(format[0], msg)
		return
	}
	span.span.RecordError(fmt.Errorf("%v", msg))
	logrus.Info(msg)
}

func Warn(span *Span, msg interface{}, format ...string) {
	if len(format) > 0 {
		span.span.RecordError(fmt.Errorf(format[0], msg))
		logrus.Warnf(format[0], msg)
		return
	}
	span.span.RecordError(fmt.Errorf("%v", msg))
	logrus.Warn(msg)
}

func Error(span *Span, msg interface{}, format ...string) {
	if len(format) > 0 {
		span.span.RecordError(fmt.Errorf(format[0], msg))
		logrus.Errorf(format[0], msg)
		return
	}
	span.span.RecordError(fmt.Errorf("%v", msg))
	logrus.Error(msg)
}

func Debug(span *Span, msg interface{}, format ...string) {
	if len(format) > 0 {
		span.span.RecordError(fmt.Errorf(format[0], msg))
		logrus.Debugf(format[0], msg)
		return
	}
	span.span.RecordError(fmt.Errorf("%v", msg))
	logrus.Debug(msg)
}
