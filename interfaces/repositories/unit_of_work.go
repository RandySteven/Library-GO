package repositories_interfaces

import (
	"context"
	"database/sql"
)

type (
	UnitOfWork interface {
		BeginTx(ctx context.Context) error
		CommitTx(ctx context.Context) error
		RollbackTx(ctx context.Context) error
		SetTx(tx *sql.Tx)
		GetTx(ctx context.Context) *sql.Tx
	}

	Trigger interface {
		PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
		ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
		QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
		QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	}
)
