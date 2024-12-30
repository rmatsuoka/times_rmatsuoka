package usercmd

import (
	"context"

	"github.com/rmatsuoka/times_rmatsuoka/internal/users"
	"github.com/rmatsuoka/times_rmatsuoka/internal/x/xsql"
)

func (c *Command) Update(ctx context.Context, db xsql.DB, code string, user *Creating) error {
	validated, err := users.ValidateCreating(user)
	if err != nil {
		return err
	}

	return db.WithinTx(ctx, func(ctx context.Context, tx xsql.Tx) error {
		id, err := c.repository().Users.ID(ctx, tx, code)
		if err != nil {
			return err
		}

		return c.repository().Users.Update(ctx, tx, id, validated)
	})
}

func Update(ctx context.Context, db xsql.DB, code string, user *Creating) error {
	return Default.Update(ctx, db, code, user)
}
