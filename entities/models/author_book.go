package models

import "time"

type AuthorBook struct {
	ID        uint64
	AuthorID  uint64
	BookID    uint64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time

	Book   Book
	Author Author
}
