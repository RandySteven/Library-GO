package repositories_interfaces

import "context"

type (
	Many2ManyRepository[T any] interface {
		Save(ctx context.Context, entity []*T) (result *T, err error)
	}

	Saver[T any] interface {
		Save(ctx context.Context, entity *T) (*T, error)
	}

	Finder[T any] interface {
		FindByID(ctx context.Context, id uint64) (*T, error)
		FindAll(ctx context.Context, skip uint64, take uint64) ([]*T, error)
	}

	Updater[T any] interface {
		Update(ctx context.Context, entity *T) (*T, error)
	}

	Deleter[T any] interface {
		DeleteByID(ctx context.Context, id uint64) error
	}
)
