package models

import (
	"github.com/RandySteven/Library-GO/enums"
	"time"
)

type EventUser struct {
	ID           uint64
	EventID      uint64
	UserID       uint64
	EventCode    string
	RedeemStatus enums.EventRedeemStatus
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}
