package models

import "time"

type User struct {
	ID             uint64
	Name           string
	Address        string
	Email          string
	PhoneNumber    string
	Password       string
	DoB            time.Time
	ProfilePicture string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time

	VerifiedAt *time.Time
}
