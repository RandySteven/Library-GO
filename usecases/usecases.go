package usecases

import (
	"github.com/RandySteven/Library-GO/caches"
	usecases_interfaces "github.com/RandySteven/Library-GO/interfaces/usecases"
	algolia_client "github.com/RandySteven/Library-GO/pkg/algolia"
	aws_client "github.com/RandySteven/Library-GO/pkg/aws"
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
}

func NewUsecases(repositories *repositories.Repositories, caches *caches.Caches, awsClient *aws_client.AWSClient, algoClient *algolia_client.AlgoliaAPISearchClient) *Usecases {
	return &Usecases{
		BagUsecase:            newBagUsecase(repositories.BagRepo, repositories.BookRepo, repositories.UserRepo),
		DevUsecase:            newDevUsecase(awsClient),
		OnboardingUsecase:     newOnboardingUsecase(repositories.UserRepo),
		BookUsecase:           newBookUsecase(repositories.UserRepo, repositories.BookRepo, repositories.GenreRepo, repositories.AuthorRepo, repositories.AuthorBookRepo, repositories.BookGenreRepo, repositories.RatingRepo, awsClient, algoClient),
		BorrowUsecase:         newBorrowUsecase(repositories.BagRepo, repositories.BookRepo, repositories.BorrowRepo, repositories.BorrowDetailRepo, repositories.UserRepo, repositories.AuthorRepo, repositories.GenreRepo),
		GenreUsecase:          newGenreUsecase(repositories.GenreRepo, repositories.BookRepo, repositories.BookGenreRepo),
		StoryGeneratorUsecase: newStoryGeneratorUsecase(awsClient),
		ReturnUsecase:         newReturnUsecase(repositories.BorrowRepo, repositories.BorrowDetailRepo, repositories.BookRepo, repositories.UserRepo),
	}
}
