package xsql

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
)

type SQLDB struct {
	DB *sql.DB
}

var _ DB = (*SQLDB)(nil)

func (s *SQLDB) QueryRow(ctx context.Context, query string, args ...any) (scan func(...any) error) {
	return queryRow(ctx, s.DB, query, args...)
}

func (s *SQLDB) Query(ctx context.Context, query string, args ...any) func(byRow func(scan func(...any))) error {
	return queryScan(ctx, s.DB, query, args...)
}

func (s *SQLDB) WithinTx(ctx context.Context, f func(context.Context, Tx) error) error {
	// skip [loggerWithPC, WithinTx]
	logger := loggerWithPC(slog.Default(), 2)
	logger.DebugContext(ctx, "begin transaction")
	sqltx, err := s.DB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		logger.ErrorContext(ctx, "failed to begin transaction", "error", err)
		return err
	}
	defer func() {
		err := sqltx.Rollback()
		// log nothing if Commit has sucessed.
		if !errors.Is(err, sql.ErrTxDone) {
			logger.DebugContext(ctx, "rollback", "error", err)
		}
	}()

	if err := f(ctx, &sqlTx{sqltx}); err != nil {
		return err
	}

	logger.DebugContext(ctx, "commit")
	if err := sqltx.Commit(); err != nil {
		logger.ErrorContext(ctx, "failed to commit transaction", "error", err)
		return err
	}
	return nil
}

type sqlTx struct {
	tx *sql.Tx
}

var _ Tx = (*sqlTx)(nil)

func (s *sqlTx) QueryRow(ctx context.Context, query string, args ...any) (scan func(...any) error) {
	return queryRow(ctx, s.tx, query, args...)
}

func (s *sqlTx) Query(ctx context.Context, query string, args ...any) func(byRow func(scan func(...any))) error {
	return queryScan(ctx, s.tx, query, args...)
}

func (s *sqlTx) Exec(ctx context.Context, query string, args ...any) error {
	// skip [loggerWithPC, Exec]
	logger := loggerWithPC(slog.Default(), 2, "query", query, "args", args)

	logger.DebugContext(ctx, "Exec")
	_, err := s.tx.ExecContext(ctx, query, args...)
	if err != nil {
		logger.ErrorContext(ctx, "failed to exec query", "error", err)
	}
	return err
}

func (s *sqlTx) ExecLastInsertID(ctx context.Context, query string, args ...any) (id int64, err error) {
	// skip [loggerWithPC, ExecLastInsertID]
	logger := loggerWithPC(slog.Default(), 2, "query", query, "args", args)

	logger.DebugContext(ctx, "ExecLastInsertID", "lastInsertedID", id)
	r, err := s.tx.ExecContext(ctx, query, args...)
	if err != nil {
		logger.ErrorContext(ctx, "failed to exec sql", "error", err)
		return 0, err
	}
	id, err = r.LastInsertId()
	if err != nil {
		slog.ErrorContext(ctx, "failed to get last insert id", "error", err)
	}
	return id, err
}

func (s *sqlTx) ExecRowsAffected(ctx context.Context, query string, args ...any) (rows int64, err error) {
	// skip [loggerWithPC, ExecRowsAffected]
	logger := loggerWithPC(slog.Default(), 2, "query", query, "args", args)

	logger.DebugContext(ctx, "ExecRowsAffected")
	r, err := s.tx.ExecContext(ctx, query, args...)
	if err != nil {
		logger.ErrorContext(ctx, "failed to exec sql", "error", err)
		return 0, err
	}
	rows, err = r.RowsAffected()
	if err != nil {
		slog.ErrorContext(ctx, "failed to get last insert id", "error", err)
	}
	return rows, err
}

type dbtx interface {
	QueryRowContext(context.Context, string, ...any) *sql.Row
	QueryContext(context.Context, string, ...any) (*sql.Rows, error)
}

var (
	_ dbtx = (*sql.DB)(nil)
	_ dbtx = (*sql.Tx)(nil)
)

func queryScan(ctx context.Context, db dbtx, query string, args ...any) func(byRow func(scan func(...any))) error {
	// skip [loggerWithPC, queryScan, DB.Query]
	logger := loggerWithPC(slog.Default(), 3, "query", query, "args", args)
	return func(byRow func(scan func(...any))) error {
		logger.Debug("Query")
		rows, err := db.QueryContext(ctx, query, args...)
		if err != nil {
			logger.ErrorContext(ctx, "failed to query", "error", err)
			return err
		}
		defer rows.Close()

		for rows.Next() {
			byRow(func(dest ...any) {
				err = rows.Scan(dest...)
			})
			if err != nil {
				logger.ErrorContext(ctx, "failed to scan", "error", err)
				return err
			}
		}
		if err = rows.Err(); err != nil {
			logger.ErrorContext(ctx, "error occured on iterating rows", "error", err)
		}
		return err
	}
}

func queryRow(ctx context.Context, db dbtx, query string, args ...any) (scan func(...any) error) {
	// skip [loggerWithPC, queryRow, DB.QueryRow]
	logger := loggerWithPC(slog.Default(), 3, "query", query, "args", args)
	return func(dest ...any) error {
		logger.DebugContext(ctx, "QueryRow")
		err := db.QueryRowContext(ctx, query, args...).Scan(dest...)
		if err != nil {
			logger.ErrorContext(ctx, "failed to query", "error", err)
		}
		return err
	}
}
