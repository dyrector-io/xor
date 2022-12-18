package ctx

import "context"

type CtxKey int

const (
	StateKey = iota
)

func SetContextVar[T any](c context.Context, key CtxKey, value T) context.Context {
	return context.WithValue(c, key, value)
}

func GetContextVar[T any](c context.Context, key CtxKey) T {
	return c.Value(key).(T)
}
