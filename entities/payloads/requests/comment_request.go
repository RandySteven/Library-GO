package requests

type (
	AddCommentRequest struct {
		BookID  uint64 `json:"book_id"`
		Comment string `json:"comment"`
	}

	ReplyCommentRequest struct {
		ParentID uint64 `json:"parent_id"`
		ReplyID  uint64 `json:"reply_id"`
		BookID   uint64 `json:"book_id"`
		Comment  string `json:"comment"`
	}

	GetCommentRequest struct {
		BookID uint64 `json:"book_id"`
		Limit  uint64 `json:"limit"`
		Page   uint64 `json:"page"`
	}
)
