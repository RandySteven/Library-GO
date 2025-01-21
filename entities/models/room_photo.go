package models

import "time"

type RoomPhoto struct {
	ID        uint64
	RoomID    uint64
	Photo     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
