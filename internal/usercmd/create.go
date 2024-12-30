package usercmd

import (
	"context"

	"github.com/rmatsuoka/times_rmatsuoka/internal/users"
	"github.com/rmatsuoka/times_rmatsuoka/internal/x/xsql"
)

type Creating struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

func (u *Creating) UserName() string { return u.Name }
func (u *Creating) UserCode() string { return u.Code }

func (c *Command) Create(ctx context.Context, db xsql.DB, user *Creating) (id users.ID, err error) {
	vuser, err := users.ValidateCreating(user)
	if err != nil {
		return nil, err
	}
	err = db.WithinTx(ctx, func(ctx context.Context, tx xsql.Tx) error {
		id, err = c.Users().Create(ctx, tx, vuser)
		return err
	})
	return id, err
}

func Create(ctx context.Context, db xsql.DB, user *Creating) (users.ID, error) {
	return Default.Create(ctx, db, user)
}
