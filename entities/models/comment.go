package models

import "time"

type Comment struct {
	ID        uint64
	UserID    uint64
	BookID    uint64
	ParentID  uint64
	Comment   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
