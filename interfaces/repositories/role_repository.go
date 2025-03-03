package repositories_interfaces

import "github.com/RandySteven/Library-GO/entities/models"

type RoleRepository interface {
	Saver[models.Role]
	Finder[models.Role]
	//UnitOfWork
}
