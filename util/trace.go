package util

import (
	"context"
)

type traceContext struct{}

func WithTraceID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, traceContext{}, id)
}

func TraceID(ctx context.Context) string {
	v, _ := ctx.Value(traceContext{}).(string)
	return v
}
