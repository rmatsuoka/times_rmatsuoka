package currnet

import (
	"context"

	"github.com/rmatsuoka/times_rmatsuoka/internal/users"
)

type key int

const userIDKey = key(0)

func UserID(ctx context.Context) users.ID {
	return ctx.Value(userIDKey).(users.ID)
}

func ContextWithUserID(ctx context.Context, id users.ID) context.Context {
	return context.WithValue(ctx, userIDKey, id)
}
