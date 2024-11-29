package clients

import (
	"context"

	"github.com/rmatsuoka/times_rmatsuoka/internal/users"
)

type key int

const clientKey = key(0)

func FromContext(ctx context.Context) users.ID {
	panic("not implement")
}

func NewContext(ctx context.Context, id users.ID) context.Context {
	return context.WithValue(ctx, clientKey, id)
}
