package util

import (
	"context"
	"fmt"
)

type traceContext struct{}

func WithTraceID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, traceContext{}, id)
}

func EnsureTraceID(ctx context.Context) context.Context {
	id := TraceID(ctx)
	if id != "" {
		return ctx
	}
	return context.WithValue(ctx, traceContext{}, fmt.Sprintf("%x", RandBytes(8)))
}

func TraceID(ctx context.Context) string {
	v, _ := ctx.Value(traceContext{}).(string)
	return v
}
