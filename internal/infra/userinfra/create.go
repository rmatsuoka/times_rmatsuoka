package userinfra

import (
	"context"

	"github.com/rmatsuoka/times_rmatsuoka/internal/infra/infratypes"
	"github.com/rmatsuoka/times_rmatsuoka/internal/users"
	"github.com/rmatsuoka/times_rmatsuoka/internal/x/xsql"
)

func (Users) Create(ctx context.Context, tx xsql.Tx, user users.ValidCreating) (users.ID, error) {
	id, err := tx.ExecLastInsertID(ctx, `
		insert into
			users (code, name)
		values
			(?, ?)
	`, user.UserCode(), user.UserName())
	return infratypes.UserID(id), err
}
