package models

import "time"

type Borrow struct {
	ID              uint64
	UserID          uint64
	BorrowReference string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time
}
