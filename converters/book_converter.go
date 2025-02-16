package converters

import (
	"github.com/RandySteven/Library-GO/entities/models"
	"github.com/RandySteven/Library-GO/entities/payloads/responses"
)

func BookToListBookResponse(b *models.Book, r *models.Rating) *responses.ListBooksResponse {
	return &responses.ListBooksResponse{
		ID:        b.ID,
		Title:     b.Title,
		Image:     b.Image,
		Status:    b.Status.ToString(),
		Rating:    r.Score,
		CreatedAt: b.CreatedAt,
		UpdatedAt: b.UpdatedAt,
		DeletedAt: b.DeletedAt,
	}
}
