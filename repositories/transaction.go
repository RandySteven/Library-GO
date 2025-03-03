package repositories

import (
	"context"
	"database/sql"
	"github.com/RandySteven/Library-GO/apperror"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	"log"
)

type (
	transaction struct {
		db *sql.DB
	}

	txCtxKey struct{}
)

func (t *transaction) RunInTx(ctx context.Context, txFunc func(ctx context.Context) *apperror.CustomError) (customErr *apperror.CustomError) {
	tx, err := t.db.BeginTx(ctx, nil)
	if err != nil {
		return apperror.NewCustomError(apperror.ErrInternalServer, `failed to begin tx`, err)
	}
	txCtx := txToContext(ctx, tx)
	if customErr = txFunc(txCtx); customErr != nil {
		_ = tx.Rollback()
		log.Println("error apa ", customErr)
		return apperror.NewCustomError(apperror.ErrInternalServer, `failed to rollback tx`, err)
	}

	if err = tx.Commit(); err != nil {
		return apperror.NewCustomError(apperror.ErrInternalServer, `failed to commit tx`, err)
	}

	return nil
}

func txToContext(ctx context.Context, tx *sql.Tx) context.Context {
	return context.WithValue(ctx, txCtxKey{}, tx)
}

func txFromContext(ctx context.Context) *sql.Tx {
	tx, ok := ctx.Value(txCtxKey{}).(*sql.Tx)
	if ok {
		return tx
	}

	return nil
}

var _ repositories_interfaces.Transaction = &transaction{}

func newTransaction(db *sql.DB) (*transaction, repositories_interfaces.DB) {
	return &transaction{
			db: db,
		}, func(ctx context.Context) repositories_interfaces.Trigger {
			if tx := txFromContext(ctx); tx != nil {
				return tx
			}
			return db
		}
}
