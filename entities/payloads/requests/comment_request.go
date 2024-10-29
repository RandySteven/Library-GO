package requests

type (
	AddCommentRequest struct {
		BookID  uint64 `json:"book_id"`
		Comment string `json:"comment"`
	}

	ReplyCommentRequest struct {
	}
)
