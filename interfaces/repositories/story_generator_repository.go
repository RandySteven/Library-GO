package repositories_interfaces

import "github.com/RandySteven/Library-GO/entities/models"

type StoryGeneratorRepository interface {
	Repository[models.StoryGenerator]
	UnitOfWork
}
