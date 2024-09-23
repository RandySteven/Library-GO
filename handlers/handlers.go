package handlers

import (
	handlers_interfaces "github.com/RandySteven/Library-GO/interfaces/handlers"
	"github.com/RandySteven/Library-GO/usecases"
)

type Handlers struct {
	OnboardingHandler handlers_interfaces.OnboardingHandler
	BookHandler       handlers_interfaces.BookHandler
	DevHandler        handlers_interfaces.DevHandler
	GenreHandler      handlers_interfaces.GenreHandler
}

func NewHandlers(usecases *usecases.Usecases) *Handlers {
	return &Handlers{
		OnboardingHandler: newOnboardingHandler(usecases.OnboardingUsecase),
		BookHandler:       newBookHandler(usecases.BookUsecase),
		DevHandler:        newDevHandler(usecases.DevUsecase),
		GenreHandler:      newGenreHandler(usecases.GenreUsecase),
	}
}
