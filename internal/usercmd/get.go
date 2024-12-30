package usercmd

import (
	"context"

	"github.com/rmatsuoka/times_rmatsuoka/internal/users"
	"github.com/rmatsuoka/times_rmatsuoka/internal/x/xsql"
)

func (c *Command) Get(ctx context.Context, db xsql.DB, code string) (*users.User, error) {
	return c.Users().GetByCode(ctx, db, code)
}

func Get(ctx context.Context, db xsql.DB, code string) (*users.User, error) {
	return Default.Get(ctx, db, code)
}
