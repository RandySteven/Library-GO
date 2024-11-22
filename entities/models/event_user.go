package models

import (
	"github.com/RandySteven/Library-GO/enums"
	"time"
)

type EventUser struct {
	ID           uint64
	UserID       uint64
	EventID      uint64
	Payed        bool
	EventCode    string
	RedeemStatus enums.EventRedeemStatus
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time

	User  User
	Event Event
}
