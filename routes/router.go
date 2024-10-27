package routes

import (
	"fmt"
	"github.com/RandySteven/Library-GO/enums"
	"github.com/RandySteven/Library-GO/handlers"
	"github.com/RandySteven/Library-GO/middlewares"
	"github.com/RandySteven/Library-GO/utils"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type (
	HandlerFunc func(w http.ResponseWriter, r *http.Request)

	Router struct {
		path    string
		handler HandlerFunc
		method  string
	}
)

func NewEndpointRouters(h *handlers.Handlers) map[enums.RouterPrefix][]*Router {
	endpointRouters := make(map[enums.RouterPrefix][]*Router)

	endpointRouters[enums.DevPrefix] = []*Router{
		Get("/health-check", h.DevHandler.HealthCheck),
		Post("/create-bucket", h.DevHandler.CreateBucket),
		Get("/buckets", h.DevHandler.GetListBuckets),
		Get("/check-points", func(writer http.ResponseWriter, request *http.Request) {
			result := []string{}
			for key, value := range endpointRouters {
				for _, r := range value {
					result = append(result, fmt.Sprintf("%s%s", key, r.path))
				}
			}
			dataKey := `endpoints`
			utils.ResponseHandler(writer, http.StatusOK, `success get list endpoint`, &dataKey, result, nil)
		}),
	}

	endpointRouters[enums.OnboardingPrefix] = []*Router{
		Post("/register", h.OnboardingHandler.RegisterUser),
		Post("/login", h.OnboardingHandler.LoginUser),
		Post("/verify", h.OnboardingHandler.VerifyUser),
	}

	endpointRouters[enums.OnboardedPrefix] = []*Router{
		Get("", h.OnboardingHandler.GetLoginUser),
	}

	endpointRouters[enums.UserPrefix] = []*Router{
		Get("/{id}", h.UserHandler.GetUserDetail),
		Get("/", h.UserHandler.GetListOfUsers),
	}

	endpointRouters[enums.BookPrefix] = []*Router{
		Post("", h.BookHandler.AddBook),
		Get("", h.BookHandler.GetAllBooks),
		Get("/{id}", h.BookHandler.GetBookByID),
		Post("/search", h.BookHandler.SearchBooks),
		Post("/rating", h.RatingHandler.SubmitRating),
	}

	endpointRouters[enums.GenrePrefix] = []*Router{
		Get("", h.GenreHandler.GetAllGenres),
		Post("", h.GenreHandler.AddNewGenre),
		Get("/{id}", h.GenreHandler.GetGenre),
	}

	endpointRouters[enums.BagPrefix] = []*Router{
		Get("", h.BagHandler.GetUserBag),
		Post("", h.BagHandler.AddBookToBag),
		Post("/remove", h.BagHandler.DeleteBookFromBag),
	}

	endpointRouters[enums.StoryPrefix] = []*Router{
		Post("", h.StoryGeneratorHandler.GenerateStory),
	}

	endpointRouters[enums.BorrowPrefix] = []*Router{
		Get("/checkout", h.BorrowHandler.BorrowCheckout),
		Get("", h.BorrowHandler.GetBorrowList),
		Get("/{id}", h.BorrowHandler.GetBorrowDetail),
		Post("/confirm", h.BorrowHandler.BorrowConfirmation),
		Post("/return", h.ReturnHandler.ReturnBook),
	}

	return endpointRouters
}

func InitRouters(routers map[enums.RouterPrefix][]*Router, r *mux.Router) {
	onboardingRouter := r.PathPrefix(enums.OnboardingPrefix.ToString()).Subrouter()
	for _, router := range routers[enums.OnboardingPrefix] {
		router.RouterLog(enums.OnboardingPrefix.ToString())
		onboardingRouter.HandleFunc(router.path, router.handler).Methods(router.method)
	}

	onboardedRouter := r.PathPrefix(enums.OnboardedPrefix.ToString()).Subrouter()
	onboardedRouter.Use(middlewares.AuthenticationMiddleware)
	for _, router := range routers[enums.OnboardedPrefix] {
		router.RouterLog(enums.OnboardedPrefix.ToString())
		onboardedRouter.HandleFunc(router.path, router.handler).Methods(router.method)
	}

	devRouter := r.PathPrefix(enums.DevPrefix.ToString()).Subrouter()
	for _, router := range routers[enums.DevPrefix] {
		router.RouterLog(enums.DevPrefix.ToString())
		devRouter.HandleFunc(router.path, router.handler).Methods(router.method)
	}

	bookRouter := r.PathPrefix(enums.BookPrefix.ToString()).Subrouter()
	bookRouter.Use(middlewares.AuthenticationMiddleware)
	for _, router := range routers[enums.BookPrefix] {
		router.RouterLog(enums.BookPrefix.ToString())
		bookRouter.HandleFunc(router.path, router.handler).Methods(router.method)
	}

	genreRouter := r.PathPrefix(enums.GenrePrefix.ToString()).Subrouter()
	for _, router := range routers[enums.GenrePrefix] {
		router.RouterLog(enums.GenrePrefix.ToString())
		genreRouter.HandleFunc(router.path, router.handler).Methods(router.method)
	}

	bagRouter := r.PathPrefix(enums.BagPrefix.ToString()).Subrouter()
	bagRouter.Use(middlewares.AuthenticationMiddleware)
	for _, router := range routers[enums.BagPrefix] {
		router.RouterLog(enums.BagPrefix.ToString())
		bagRouter.HandleFunc(router.path, router.handler).Methods(router.method)
	}

	storyRouter := r.PathPrefix(enums.StoryPrefix.ToString()).Subrouter()
	for _, router := range routers[enums.StoryPrefix] {
		router.RouterLog(enums.StoryPrefix.ToString())
		storyRouter.HandleFunc(router.path, router.handler).Methods(router.method)
	}

	borrowRouter := r.PathPrefix(enums.BorrowPrefix.ToString()).Subrouter()
	borrowRouter.Use(middlewares.AuthenticationMiddleware)
	for _, router := range routers[enums.BorrowPrefix] {
		router.RouterLog(enums.BorrowPrefix.ToString())
		borrowRouter.HandleFunc(router.path, router.handler).Methods(router.method)
	}

	userRouter := r.PathPrefix(enums.UserPrefix.ToString()).Subrouter()
	userRouter.Use(middlewares.AuthenticationMiddleware)
	for _, router := range routers[enums.UserPrefix] {
		router.RouterLog(enums.UserPrefix.ToString())
		userRouter.HandleFunc(router.path, router.handler).Methods(router.method)
	}
}

func (router *Router) RouterLog(prefix string) {
	log.Printf("%12s | %4s/ \n", router.method, prefix+router.path)
}
