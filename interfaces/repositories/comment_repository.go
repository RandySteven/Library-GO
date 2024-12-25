package repositories_interfaces

import (
	"context"
	"github.com/RandySteven/Library-GO/entities/models"
)

type CommentRepository interface {
	Saver[models.Comment]
	Finder[models.Comment]
	Deleter[models.Comment]
	UnitOfWork
	FindCommentsByBookID(ctx context.Context, bookID uint64) (result []*models.Comment, err error)
}
