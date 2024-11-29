package userinfra

import (
	"context"

	"github.com/rmatsuoka/times_rmatsuoka/internal/users"
	"github.com/rmatsuoka/times_rmatsuoka/internal/x/xsql"
)

func (User) Get(ctx context.Context, db xsql.Querier, id users.ID) (*users.User, error) {
	u := users.User{}
	err := db.QueryRow(ctx, `
		select
			id, code, name, created_at, updated_at
		from
			users
		where
			id = ?
	`, id.(ID))(&u.ID, &u.Code, &u.Name, &u.CreatedAt, &u.UpdatedAt)
	return &u, err
}

func (User) ID(ctx context.Context, db xsql.Querier, code string) (users.ID, error) {
	panic("not implements")
}

func (User) GetByCode(ctx context.Context, db xsql.Querier, code string) (*users.User, error) {
	panic("not implements")
}
