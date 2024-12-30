package userinfra

import (
	"context"

	"github.com/rmatsuoka/times_rmatsuoka/internal/infra/infratypes"
	"github.com/rmatsuoka/times_rmatsuoka/internal/users"
	"github.com/rmatsuoka/times_rmatsuoka/internal/x/xsql"
)

func (u Users) Update(ctx context.Context, tx xsql.Tx, id users.ID, user users.ValidCreating) error {
	current, err := u.Get(ctx, tx, id)
	if err != nil {
		return infratypes.WrapError(err)
	}

	if current.Code != user.UserCode() {
		err := tx.Exec(ctx, `insert into usercodes (code) values (?)`, user.UserCode())
		if err != nil {
			return infratypes.WrapError(err)
		}
	}

	err = tx.Exec(ctx, `
update
	users
set
	code = ?,
	name = ?
where
	id = ?
`, user.UserCode(), user.UserName(), id)
	return infratypes.WrapError(err)
}
