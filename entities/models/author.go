package models

import "time"

type Author struct {
	ID          uint64
	Name        string
	Nationality string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}
