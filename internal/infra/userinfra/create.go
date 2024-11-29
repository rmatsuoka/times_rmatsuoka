package userinfra

import (
	"context"

	"github.com/rmatsuoka/times_rmatsuoka/internal/users"
	"github.com/rmatsuoka/times_rmatsuoka/internal/x/xsql"
)

func (User) Create(ctx context.Context, tx xsql.Tx, user users.ValidCreating) (users.ID, error) {
	panic("not implement")
}
