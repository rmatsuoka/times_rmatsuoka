package userinfra

import (
	"context"

	"github.com/rmatsuoka/times_rmatsuoka/internal/infra/infratypes"
	"github.com/rmatsuoka/times_rmatsuoka/internal/users"
	"github.com/rmatsuoka/times_rmatsuoka/internal/x/xsql"
)

func (Users) Delete(ctx context.Context, tx xsql.Tx, id users.ID) error {
	err := tx.Exec(ctx, `
insert into
	deleted_users (user_id, code, name, created_at, updated_at)
select
	id, code, name, created_at, updated_at
from
	users
where
	id = ?`, id)
	if err != nil {
		return infratypes.WrapError(err)
	}

	err = tx.Exec(ctx, `delete from users where id = ?`, id)
	return infratypes.WrapError(err)
}
