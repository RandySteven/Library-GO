package repositories_interfaces

import (
	"context"
	"database/sql"
	"github.com/RandySteven/Library-GO/apperror"
)

type (
	DB func(ctx context.Context) Trigger

	Trigger interface {
		PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
		ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
		QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
		QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	}

	Transaction interface {
		RunInTx(ctx context.Context, txFnc func(ctx context.Context) *apperror.CustomError) (err *apperror.CustomError)
	}
)
