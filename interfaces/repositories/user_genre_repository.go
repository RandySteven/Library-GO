package repositories_interfaces

import "github.com/RandySteven/Library-GO/entities/models"

type UserGenreRepository interface {
	Saver[models.UserGenre]
	Finder[models.UserGenre]
}
