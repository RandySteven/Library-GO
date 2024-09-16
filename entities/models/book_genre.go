package models

import "time"

type BookGenre struct {
	ID        uint64
	BookID    uint64
	GenreID   uint64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
