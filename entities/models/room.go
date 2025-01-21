package models

import "time"

type Room struct {
	ID          uint64
	Name        string
	IsAvailable bool
	Thumbnail   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}
