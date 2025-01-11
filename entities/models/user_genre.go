package models

type UserGenre struct {
	ID      uint64
	UserID  uint64
	GenreID uint64

	User  *User
	Genre *Genre
}
