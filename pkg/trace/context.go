package trace

import (
	"context"
	"go.uber.org/zap"
)

type contextKey struct {
}

var (
	traceKey = &contextKey{}
)

func SetFields(ctx context.Context, traceID, spanID string) context.Context {
	return context.WithValue(ctx, traceKey, []zap.Field{
		zap.String("trace_id", traceID),
		zap.String("span_id", spanID),
	})
}

func GetFields(ctx context.Context) []zap.Field {
	value, ok := ctx.Value(traceKey).([]zap.Field)
	if ok {
		return value
	}
	return nil
}
