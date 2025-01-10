package messages

type BorrowMessage struct {
	ID              string `json:"id"`
	BorrowReference string `json:"borrow_reference"`
}
