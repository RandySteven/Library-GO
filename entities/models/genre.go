package models

import "time"

type Genre struct {
	ID        uint64
	Genre     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
