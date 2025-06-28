package scenario

import (
	"context"
)

// ContextKey is the key used to store scenario context in context.Context
type ContextKey struct{}

func WithContext(ctx context.Context, scenarioCtx *Context) context.Context {
	return context.WithValue(ctx, ContextKey{}, scenarioCtx)
}

func FromContext(ctx context.Context) *Context {
	if scenarioCtx, ok := ctx.Value(ContextKey{}).(*Context); ok {
		return scenarioCtx
	}
	return nil
}

func MustFromContext(ctx context.Context) *Context {
	scenarioCtx := FromContext(ctx)
	if scenarioCtx == nil {
		panic("scenario context not found in context")
	}
	return scenarioCtx
}
