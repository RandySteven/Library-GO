package models

import "time"

type Rating struct {
	ID        uint64
	BookID    uint64
	UserID    uint64
	Score     float32
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time

	Book Book
}
