package models

type (
	Bag struct {
		ID     uint64
		UserID uint64
		BookID uint64
	}

	BookBagCache struct {
		BookID string
		UserID string
		Status string
	}
)
