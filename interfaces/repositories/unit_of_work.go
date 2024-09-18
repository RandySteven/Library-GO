package repositories_interfaces

import (
	"context"
	"database/sql"
)

type UnitOfWork interface {
	BeginTx(ctx context.Context) error
	CommitTx(ctx context.Context) error
	RollbackTx(ctx context.Context) error
	SetTx(tx *sql.Tx)
	GetTx(ctx context.Context) *sql.Tx
}
