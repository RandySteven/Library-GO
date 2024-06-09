package repositories_interfaces

import (
	"github.com/RandySteven/Library-GO/repositories"
	"github.com/RandySteven/Library-GO/user_service/entities/models"
)

type UserRepository interface {
	repositories.Repository[models.User]
}
