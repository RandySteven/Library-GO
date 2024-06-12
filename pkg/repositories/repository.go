package repositories

import "context"

type Repository[T any] interface {
	Create(ctx context.Context, models *T) (uint64, error) //return the id
	FindByID(ctx context.Context, id uint64) (*T, error)   //return specific data and error
	FindAll(ctx context.Context) ([]*T, error)             //return array of models and error
	Update(ctx context.Context, models *T) (*T, error)     //return update of model and error
	Delete(ctx context.Context, id uint64) error           //return delete error
}
