package repository

import (
	"context"
	"iter"

	"github.com/rmatsuoka/times_rmatsuoka/internal/users"
	"github.com/rmatsuoka/times_rmatsuoka/internal/x/xsql"
)

type Users interface {
	ID(ctx context.Context, db xsql.Querier, code string) (users.ID, error)
	Get(ctx context.Context, db xsql.Querier, id users.ID) (*users.User, error)
	GetByCode(ctx context.Context, db xsql.Querier, code string) (*users.User, error)
	GetMany(ctx context.Context, db xsql.Querier, ids iter.Seq[users.ID]) (map[users.ID]*users.User, error)
	Create(ctx context.Context, tx xsql.Tx, user users.ValidCreating) (users.ID, error)

	Update(ctx context.Context, tx xsql.Tx, id users.ID, user users.ValidCreating) error
	Delete(ctx context.Context, tx xsql.Tx, id users.ID) error
}
