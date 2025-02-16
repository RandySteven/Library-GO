package models

import (
	"github.com/RandySteven/Library-GO/enums"
	"time"
)

type Book struct {
	ID          uint64
	Title       string
	Description string
	Image       string
	Status      enums.BookStatus
	PDFFile     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}
