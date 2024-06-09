package models

import "time"

type (
	User struct {
		ID          uint64
		FullName    string
		UserName    string
		IsVerify    bool
		DateOfBirth time.Time
		CreatedAt   time.Time
		CreatedBy   uint64
		UpdatedAt   time.Time
		UpdatedBy   uint64
		DeletedAt   *time.Time
		DeletedBy   uint64
	}

	Profile struct {
		ID        uint64
		Email     string
		Password  string
		UserID    uint64
		CreatedAt time.Time
		CreatedBy uint64
		UpdatedAt time.Time
		UpdatedBy uint64
		DeletedAt *time.Time
		DeletedBy uint64
	}

	Role struct {
		ID   uint64
		Role string
	}
)
