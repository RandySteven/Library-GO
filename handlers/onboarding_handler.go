package handlers

import (
	handlers_interfaces "github.com/RandySteven/Library-GO/interfaces/handlers"
	usecases_interfaces "github.com/RandySteven/Library-GO/interfaces/usecases"
	"net/http"
)

type OnboardingHandler struct {
	onboardingUsecase usecases_interfaces.OnboardingUsecase
}

func (o *OnboardingHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
}

func (o *OnboardingHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
}

func (o *OnboardingHandler) VerifyUser(w http.ResponseWriter, r *http.Request) {
}

var _ handlers_interfaces.OnboardingHandler = &OnboardingHandler{}

func newOnboardingHandler(onboardingUsecase usecases_interfaces.OnboardingUsecase) *OnboardingHandler {
	return &OnboardingHandler{
		onboardingUsecase: onboardingUsecase,
	}
}
