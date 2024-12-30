package userinfra

import (
	"context"
	"iter"

	"github.com/rmatsuoka/times_rmatsuoka/internal/infra/infratypes"
	"github.com/rmatsuoka/times_rmatsuoka/internal/infra/schema"
	"github.com/rmatsuoka/times_rmatsuoka/internal/users"
	"github.com/rmatsuoka/times_rmatsuoka/internal/x/xsql"
)

func (Users) Get(ctx context.Context, db xsql.Querier, id users.ID) (*users.User, error) {
	var u schema.User
	err := db.QueryRow(ctx, `
		select
			id, code, name, created_at, updated_at
		from
			users
		where
			id = ?
	`, id.(infratypes.UserID))(&u.ID, &u.Code, &u.Name, &u.CreatedAt, &u.UpdatedAt)
	return infratypes.UsersUser(&u), infratypes.WrapError(err)
}

func (Users) ID(ctx context.Context, db xsql.Querier, code string) (users.ID, error) {
	var id infratypes.UserID
	err := db.QueryRow(ctx, `select id from users where code = ?`, code)(&id)
	return id, infratypes.WrapError(err)
}

func (Users) GetByCode(ctx context.Context, db xsql.Querier, code string) (*users.User, error) {
	var u schema.User
	err := db.QueryRow(ctx, `
		select
			id, code, name, created_at, updated_at
		from
			users
		where
			code = ?
	`, code)(&u.ID, &u.Code, &u.Name, &u.CreatedAt, &u.UpdatedAt)
	return infratypes.UsersUser(&u), infratypes.WrapError(err)
}

func (Users) GetMany(ctx context.Context, db xsql.Querier, ids iter.Seq[users.ID]) (map[users.ID]*users.User, error) {
	slice := xsql.CollectAny(ids)
	if len(slice) == 0 {
		return nil, nil
	}

	_ = slice[0].(infratypes.UserID) // assert that type of userid is infratypes.UserID.

	q := xsql.ListQuery(`
		select
			id, code, name, created_at, updated_at
		from
			users
		where
			id in ({{?}})
	`, len(slice))

	usermap := make(map[users.ID]*users.User, len(slice))
	err := db.Query(ctx, q, slice...)(func(scan func(...any)) {
		var u schema.User
		scan(&u.ID, &u.Code, &u.Name, &u.CreatedAt, &u.UpdatedAt)
		usermap[infratypes.UserID(u.ID)] = infratypes.UsersUser(&u)
	})
	return usermap, infratypes.WrapError(err)
}
