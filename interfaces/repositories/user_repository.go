package repositories_interfaces

import "github.com/RandySteven/Library-GO/entities/models"

type UserRepository interface {
	Repository[models.User]
}
