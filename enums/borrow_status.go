package enums

type BorrowStatus int

const (
	Processing BorrowStatus = iota + 1
	Success
	Failed
)
