package opentelemetry

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

func InitSpan(ctx context.Context, tracerName string, operation string) trace.Span {
	tracer := otel.Tracer(tracerName)
	_, span := tracer.Start(ctx, operation)
	return span
}
