package models

import "time"

type RoomChat struct {
	ID        uint64
	RoomName  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
