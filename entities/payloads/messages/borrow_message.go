package messages

import "time"

type (
	BorrowMessage struct {
		ID              string         `json:"id"`
		BorrowReference string         `json:"borrow_reference"`
		BorrowBooks     []*BorrowBooks `json:"books"`
	}

	BorrowBooks struct {
		ID        uint64     `json:"id"`
		Title     string     `json:"title"`
		Image     string     `json:"image"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt time.Time  `json:"updated_at"`
		DeletedAt *time.Time `json:"deleted_at"`
	}
)
