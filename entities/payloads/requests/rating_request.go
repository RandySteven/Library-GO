package requests

type RatingRequest struct {
	BookID uint64  `json:"book_id"`
	Rating float32 `json:"rating"`
}
