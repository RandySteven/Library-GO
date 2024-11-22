package models

import "time"

type BorrowDetail struct {
	ID                 uint64
	BorrowID           uint64
	BookID             uint64
	BorrowedDate       time.Time
	ReturnedDate       time.Time
	VerifiedReturnDate *time.Time
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          *time.Time

	Borrow Borrow
	Book   Book
}
