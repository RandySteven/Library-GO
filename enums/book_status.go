package enums

type BookStatus int

const (
	Available BookStatus = iota + 1
	Borrowed
)
