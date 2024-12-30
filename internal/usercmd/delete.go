package usercmd

import (
	"context"

	"github.com/rmatsuoka/times_rmatsuoka/internal/x/xsql"
)

func (c *Command) Delete(ctx context.Context, db xsql.DB, code string) error {
	return db.WithinTx(ctx, func(ctx context.Context, tx xsql.Tx) error {
		id, err := c.Users().ID(ctx, tx, code)
		if err != nil {
			return err
		}
		return c.Users().Delete(ctx, tx, id)
	})
}

func Delete(ctx context.Context, db xsql.DB, code string) error {
	return Default.Delete(ctx, db, code)
}
