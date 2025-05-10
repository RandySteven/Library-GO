package usecases

import (
	"github.com/RandySteven/Library-GO/caches"
	usecases_interfaces "github.com/RandySteven/Library-GO/interfaces/usecases"
	algolia_client "github.com/RandySteven/Library-GO/pkg/algolia"
	aws_client "github.com/RandySteven/Library-GO/pkg/aws"
	oauth2_client "github.com/RandySteven/Library-GO/pkg/oauth2"
	rabbitmqs_client "github.com/RandySteven/Library-GO/pkg/rabbitmqs"
	"github.com/RandySteven/Library-GO/repositories"
)

type Usecases struct {
	BookUsecase           usecases_interfaces.BookUsecase
	BorrowUsecase         usecases_interfaces.BorrowUsecase
	OnboardingUsecase     usecases_interfaces.OnboardingUsecase
	BagUsecase            usecases_interfaces.BagUsecase
	UserUsecase           usecases_interfaces.UserUsecase
	DevUsecase            usecases_interfaces.DevUsecase
	GenreUsecase          usecases_interfaces.GenreUsecase
	StoryGeneratorUsecase usecases_interfaces.StoryGeneratorUsecase
	ReturnUsecase         usecases_interfaces.ReturnUsecase
	RatingUsecase         usecases_interfaces.RatingUsecase
	CommentUsecase        usecases_interfaces.CommentUsecase
	EventUsecase          usecases_interfaces.EventUsecase
	ChatUsecase           usecases_interfaces.ChatUsecase
	RoomUsecase           usecases_interfaces.RoomUsecase
}

func NewUsecases(repositories *repositories.Repositories, caches *caches.Caches, awsClient aws_client.AWS, algoClient algolia_client.AlgoliaAPISearch, pubsub rabbitmqs_client.PubSub, oauth2 oauth2_client.Oauth2) *Usecases {
	return &Usecases{
		BagUsecase:            newBagUsecase(repositories.BagRepo, repositories.BookRepo, repositories.UserRepo, caches.BagCache, repositories.Transaction),
		DevUsecase:            newDevUsecase(awsClient, pubsub),
		OnboardingUsecase:     newOnboardingUsecase(repositories.UserRepo, repositories.RoleUserRepo, pubsub, repositories.Transaction, oauth2),
		BookUsecase:           newBookUsecase(repositories.UserRepo, repositories.BookRepo, repositories.GenreRepo, repositories.AuthorRepo, repositories.AuthorBookRepo, repositories.BookGenreRepo, repositories.BorrowRepo, repositories.BorrowDetailRepo, repositories.RatingRepo, awsClient, algoClient, caches.BookCache, pubsub, repositories.Transaction),
		BorrowUsecase:         newBorrowUsecase(repositories.BagRepo, repositories.BookRepo, repositories.BorrowRepo, repositories.BorrowDetailRepo, repositories.UserRepo, repositories.AuthorRepo, repositories.GenreRepo, caches.BorrowCache, caches.BookCache, pubsub, repositories.Transaction),
		GenreUsecase:          newGenreUsecase(repositories.GenreRepo, repositories.BookRepo, repositories.BookGenreRepo, repositories.RatingRepo, caches.GenreCache),
		StoryGeneratorUsecase: newStoryGeneratorUsecase(awsClient),
		ReturnUsecase:         newReturnUsecase(repositories.BorrowRepo, repositories.BorrowDetailRepo, repositories.BookRepo, repositories.UserRepo, repositories.Transaction),
		RatingUsecase:         newRatingUsecase(repositories.RatingRepo),
		UserUsecase:           newUserUsecase(repositories.UserRepo),
		CommentUsecase:        newCommentUsecase(repositories.CommentRepo, repositories.UserRepo, repositories.BookRepo),
		EventUsecase:          newEventUsecase(awsClient, repositories.EventRepo, repositories.EventUserRepo, repositories.UserRepo),
		ChatUsecase:           newChatUsecase(repositories.ChatRepo, repositories.RoomChatRepo, repositories.RoomChatUserRepo, repositories.UserRepo),
		RoomUsecase:           newRoomUsecase(repositories.RoomRepo, repositories.RoomPhotoRepo, awsClient),
	}
}
