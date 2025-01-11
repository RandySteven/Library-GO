package responses

import "time"

type (
	CommentResponse struct {
		ID      uint64 `json:"id"`
		UserID  uint64 `json:"user_id"`
		BookID  uint64 `json:"book_id"`
		Comment string `json:"comment"`
	}

	ReplyCommentResponse struct {
		ID        uint64 `json:"id"`
		UserID    uint64 `json:"user_id"`
		BookID    uint64 `json:"book_id"`
		CommentID uint64 `json:"parent_id"`
		Comment   string `json:"comment"`
	}

	ListBookCommentsResponse struct {
		ID       uint64  `json:"id"`
		BookID   uint64  `json:"book_id"`
		ParentID *uint64 `json:"parent_id"`
		Comment  string  `json:"comment"`
		User     struct {
			ID    uint64 `json:"id"`
			Name  string `json:"name"`
			Email string `json:"consumers"`
		} `json:"user"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
