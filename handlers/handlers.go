package handlers

import (
	handlers_interfaces "github.com/RandySteven/Library-GO/interfaces/handlers"
	"github.com/RandySteven/Library-GO/usecases"
)

type Handlers struct {
	OnboardingHandler handlers_interfaces.OnboardingHandler
}

func NewHandlers(usecases *usecases.Usecases) *Handlers {
	return &Handlers{
		OnboardingHandler: newOnboardingHandler(usecases.OnboardingUsecase),
	}
}
