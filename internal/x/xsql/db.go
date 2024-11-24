package xsql

import "context"

type DB interface {
	Querier
	WithinTx(context.Context, func(context.Context, Tx) error) error
}

type Tx interface {
	Querier
	Exec(ctx context.Context, query string, args ...any) error
	ExecLastInsertID(ctx context.Context, query string, args ...any) (id int64, err error)
	ExecRowsAffected(ctx context.Context, query string, args ...any) (rows int64, err error)
}

type Querier interface {
	Query(ctx context.Context, query string, args ...any) func(byRow func(scan func(...any))) error
	QueryRow(ctx context.Context, query string, args ...any) (scan func(...any) error)
}
