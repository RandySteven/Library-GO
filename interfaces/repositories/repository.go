package repositories_interfaces

import "context"

type (
	Repository[T any] interface {
		Save(ctx context.Context, entity *T) (result *T, err error)
		FindByID(ctx context.Context, id uint64) (result *T, err error)
		FindAll(ctx context.Context, skip uint64, take uint64) ([]*T, error)
		DeleteByID(ctx context.Context, id uint64) (err error)
		Update(ctx context.Context, entity *T) (result *T, err error)
	}

	Many2ManyRepository[T any] interface {
		Save(ctx context.Context, entity []*T) (result *T, err error)
	}
)
