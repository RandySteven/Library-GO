package usecases

import (
	"github.com/RandySteven/Library-GO/caches"
	usecases_interfaces "github.com/RandySteven/Library-GO/interfaces/usecases"
	"github.com/RandySteven/Library-GO/repositories"
)

type Usecases struct {
	BookUsecase       usecases_interfaces.BookUsecase
	BorrowUsecase     usecases_interfaces.BorrowUsecase
	OnboardingUsecase usecases_interfaces.OnboardingUsecase
	UserUsecase       usecases_interfaces.UserUsecase
}

func NewUsecases(repositories *repositories.Repositories, caches *caches.Caches) *Usecases {
	return &Usecases{
		OnboardingUsecase: newOnboardingUsecase(repositories.UserRepo),
		BookUsecase:       newBookUsecase(repositories.UserRepo, repositories.BookRepo, repositories.GenreRepo, repositories.AuthorRepo, repositories.AuthorBookRepo, repositories.BookGenreRepo),
	}
}
