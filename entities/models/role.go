package models

import "time"

type Role struct {
	ID        uint64
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
