package usecases

import (
	usecases_interfaces "github.com/RandySteven/Library-GO/interfaces/usecases"
	"github.com/RandySteven/Library-GO/repositories"
)

type Usecases struct {
	BookUsecase       usecases_interfaces.BookUsecase
	BorrowUsecase     usecases_interfaces.BorrowUsecase
	OnboardingUsecase usecases_interfaces.OnboardingUsecase
}

func NewUsecases(repositories *repositories.Repositories) *Usecases {
	return &Usecases{}
}
