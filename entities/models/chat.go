package models

type Chat struct {
	ID         uint64
	RoomChatID uint64
	UserID     uint64
	Chat       string
}
