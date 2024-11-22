package models

import "time"

type RoomChatUser struct {
	ID         uint64
	RoomChatID uint64
	UserID     uint64
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time

	RoomChat RoomChat
	User     User
}
