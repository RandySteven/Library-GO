package handlers

import (
	handlers_interfaces "github.com/RandySteven/Library-GO/interfaces/handlers"
	"github.com/RandySteven/Library-GO/usecases"
)

type Handlers struct {
	OnboardingHandler     handlers_interfaces.OnboardingHandler
	BookHandler           handlers_interfaces.BookHandler
	DevHandler            handlers_interfaces.DevHandler
	GenreHandler          handlers_interfaces.GenreHandler
	BagHandler            handlers_interfaces.BagHandler
	StoryGeneratorHandler handlers_interfaces.StoryGeneratorHandler
	BorrowHandler         handlers_interfaces.BorrowHandler
	ReturnHandler         handlers_interfaces.ReturnHandler
	RatingHandler         handlers_interfaces.RatingHandler
	UserHandler           handlers_interfaces.UserHandler
}

func NewHandlers(usecases *usecases.Usecases) *Handlers {
	return &Handlers{
		OnboardingHandler:     newOnboardingHandler(usecases.OnboardingUsecase),
		BookHandler:           newBookHandler(usecases.BookUsecase),
		DevHandler:            newDevHandler(usecases.DevUsecase),
		GenreHandler:          newGenreHandler(usecases.GenreUsecase),
		BagHandler:            newBagHandler(usecases.BagUsecase),
		StoryGeneratorHandler: newStoryGeneratorHandler(usecases.StoryGeneratorUsecase),
		BorrowHandler:         newBorrowHandler(usecases.BorrowUsecase),
		ReturnHandler:         newReturnHandler(usecases.ReturnUsecase),
		RatingHandler:         newRatingHandler(usecases.RatingUsecase),
		UserHandler:           newUserHandler(usecases.UserUsecase),
	}
}
