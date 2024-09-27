package usecases_interfaces

import "context"

type ReturnUsecase interface {
	ReturnBook(ctx context.Context)
}
