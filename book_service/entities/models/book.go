package models

import "time"

type Book struct {
	ID          int64
	Title       string
	Description string
	Author      string
	Image       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}
