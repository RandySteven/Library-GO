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
		path        string
		handler     HandlerFunc
		method      string
		middlewares []enums.Middleware
	}

	RouterPrefix map[enums.RouterPrefix][]*Router
)

func NewEndpointRouters(h *handlers.Handlers) RouterPrefix {
	endpointRouters := make(RouterPrefix)

	endpointRouters[enums.DevPrefix] = []*Router{
		Get("/health-check", h.DevHandler.HealthCheck, enums.RateLimiterMiddleware),
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
		}, enums.RateLimiterMiddleware),
	}

	endpointRouters[enums.OnboardingPrefix] = []*Router{
		Post("/register", h.OnboardingHandler.RegisterUser),
		Post("/login", h.OnboardingHandler.LoginUser, enums.RateLimiterMiddleware),
		Post("/verify", h.OnboardingHandler.VerifyUser),
	}

	endpointRouters[enums.OnboardedPrefix] = []*Router{
		Get("", h.OnboardingHandler.GetLoginUser),
	}

	endpointRouters[enums.UserPrefix] = []*Router{
		Get("/{id}", h.UserHandler.GetUserDetail, enums.RateLimiterMiddleware, enums.AuthenticationMiddleware),
		Get("", h.UserHandler.GetListOfUsers, enums.RateLimiterMiddleware, enums.AuthenticationMiddleware),
	}

	endpointRouters[enums.BookPrefix] = []*Router{
		Get("", h.BookHandler.GetAllBooks),
		Post("", h.BookHandler.AddBook, enums.RateLimiterMiddleware, enums.AuthenticationMiddleware),
		Get("/{id}", h.BookHandler.GetBookByID, enums.RateLimiterMiddleware),
		Post("/search", h.BookHandler.SearchBooks),
	}

	endpointRouters[enums.RatingPrefix] = []*Router{
		Post("", h.RatingHandler.SubmitRating, enums.RateLimiterMiddleware, enums.AuthenticationMiddleware),
		Post("/sort", h.RatingHandler.BookOrdersRating),
	}

	endpointRouters[enums.GenrePrefix] = []*Router{
		Get("", h.GenreHandler.GetAllGenres),
		Post("", h.GenreHandler.AddNewGenre, enums.RateLimiterMiddleware, enums.AuthenticationMiddleware),
		Get("/{id}", h.GenreHandler.GetGenre, enums.RateLimiterMiddleware),
	}

	endpointRouters[enums.BagPrefix] = []*Router{
		Get("", h.BagHandler.GetUserBag, enums.AuthenticationMiddleware),
		Post("", h.BagHandler.AddBookToBag, enums.RateLimiterMiddleware, enums.AuthenticationMiddleware),
		Post("/remove", h.BagHandler.DeleteBookFromBag, enums.AuthenticationMiddleware),
	}

	endpointRouters[enums.StoryPrefix] = []*Router{
		Post("", h.StoryGeneratorHandler.GenerateStory),
	}

	endpointRouters[enums.BorrowPrefix] = []*Router{
		Get("/checkout", h.BorrowHandler.BorrowCheckout, enums.AuthenticationMiddleware),
		Get("", h.BorrowHandler.GetBorrowList, enums.AuthenticationMiddleware, enums.RateLimiterMiddleware),
		Get("/{id}", h.BorrowHandler.GetBorrowDetail, enums.AuthenticationMiddleware),
		Post("/confirm", h.BorrowHandler.BorrowConfirmation),
		Post("/return", h.ReturnHandler.ReturnBook),
	}

	endpointRouters[enums.CommentPrefix] = []*Router{
		Post("", h.CommentHandler.Comment, enums.AuthenticationMiddleware),
		Post("/reply", h.CommentHandler.Reply, enums.AuthenticationMiddleware),
		Post("/get", h.CommentHandler.GetBookComment, enums.RateLimiterMiddleware),
	}

	return endpointRouters
}

func InitRouters(routers RouterPrefix, r *mux.Router) {
	whitelistedMiddleware := middlewares.NewWhitelistedMiddleware()
	middlewareValidator := middlewares.NewMiddlewareValidator(whitelistedMiddleware)

	r.Use(
		middlewares.LoggingMiddleware,
		middlewares.CorsMiddleware,
		middlewares.TimeoutMiddleware,
		middlewareValidator.RateLimiterMiddleware,
	)

	onboardingRouter := r.PathPrefix(enums.OnboardingPrefix.ToString()).Subrouter()
	for _, router := range routers[enums.OnboardingPrefix] {
		whitelistedMiddleware.RegisterMiddleware(enums.OnboardedPrefix, router.method, router.path, router.middlewares)
		router.RouterLog(enums.OnboardingPrefix.ToString())
		onboardingRouter.HandleFunc(router.path, router.handler).Methods(router.method)
	}

	onboardedRouter := r.PathPrefix(enums.OnboardedPrefix.ToString()).Subrouter()
	onboardedRouter.Use(middlewareValidator.AuthenticationMiddleware)
	for _, router := range routers[enums.OnboardedPrefix] {
		whitelistedMiddleware.RegisterMiddleware(enums.OnboardedPrefix, router.method, router.path, router.middlewares)
		router.RouterLog(enums.OnboardedPrefix.ToString())
		onboardedRouter.HandleFunc(router.path, router.handler).Methods(router.method)
	}

	devRouter := r.PathPrefix(enums.DevPrefix.ToString()).Subrouter()
	for _, router := range routers[enums.DevPrefix] {
		whitelistedMiddleware.RegisterMiddleware(enums.DevPrefix, router.method, router.path, router.middlewares)
		router.RouterLog(enums.DevPrefix.ToString())
		devRouter.HandleFunc(router.path, router.handler).Methods(router.method)
	}

	bookRouter := r.PathPrefix(enums.BookPrefix.ToString()).Subrouter()
	bookRouter.Use(middlewareValidator.AuthenticationMiddleware)
	for _, router := range routers[enums.BookPrefix] {
		whitelistedMiddleware.RegisterMiddleware(enums.BookPrefix, router.method, router.path, router.middlewares)
		router.RouterLog(enums.BookPrefix.ToString())
		bookRouter.HandleFunc(router.path, router.handler).Methods(router.method)
	}

	genreRouter := r.PathPrefix(enums.GenrePrefix.ToString()).Subrouter()
	genreRouter.Use(middlewareValidator.AuthenticationMiddleware)
	for _, router := range routers[enums.GenrePrefix] {
		whitelistedMiddleware.RegisterMiddleware(enums.GenrePrefix, router.method, router.path, router.middlewares)
		router.RouterLog(enums.GenrePrefix.ToString())
		genreRouter.HandleFunc(router.path, router.handler).Methods(router.method)
	}

	bagRouter := r.PathPrefix(enums.BagPrefix.ToString()).Subrouter()
	bagRouter.Use(middlewareValidator.AuthenticationMiddleware)
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
	borrowRouter.Use(middlewareValidator.AuthenticationMiddleware)
	for _, router := range routers[enums.BorrowPrefix] {
		router.RouterLog(enums.BorrowPrefix.ToString())
		borrowRouter.HandleFunc(router.path, router.handler).Methods(router.method)
	}

	userRouter := r.PathPrefix(enums.UserPrefix.ToString()).Subrouter()
	userRouter.Use(middlewareValidator.AuthenticationMiddleware)
	for _, router := range routers[enums.UserPrefix] {
		router.RouterLog(enums.UserPrefix.ToString())
		userRouter.HandleFunc(router.path, router.handler).Methods(router.method)
	}

	commentRouter := r.PathPrefix(enums.CommentPrefix.ToString()).Subrouter()
	commentRouter.Use(middlewareValidator.AuthenticationMiddleware)
	for _, router := range routers[enums.CommentPrefix] {
		router.RouterLog(enums.CommentPrefix.ToString())
		commentRouter.HandleFunc(router.path, router.handler).Methods(router.method)
	}

	ratingRouter := r.PathPrefix(enums.RatingPrefix.ToString()).Subrouter()
	for _, router := range routers[enums.RatingPrefix] {
		router.RouterLog(enums.RatingPrefix.ToString())
		ratingRouter.HandleFunc(router.path, router.handler).Methods(router.method)
	}
}

func (router *Router) RouterLog(prefix string) {
	log.Printf("%12s | %4s/ \n", router.method, prefix+router.path)
}
