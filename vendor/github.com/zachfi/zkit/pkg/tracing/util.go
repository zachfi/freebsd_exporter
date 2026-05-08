package tracing

import (
	"fmt"
	"log/slog"

	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func ErrHandler(span trace.Span, err error, message string, l *slog.Logger) error {
	defer span.End()

	if err != nil {
		if l != nil {
			l.Error(message, "err", err)
		}
		span.SetStatus(codes.Error, fmt.Errorf("%s: %w", message, err).Error())
	} else {
		span.SetStatus(codes.Ok, "ok")
	}
	return err
}
