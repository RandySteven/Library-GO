package enums

type BookStatus int

const (
	Available BookStatus = iota + 1
	Borrowed
)

func (b BookStatus) ToString() string {
	switch b {
	case Available:
		return "Available"
	case Borrowed:
		return "Borrowed"
	}
	return "Unknown"
}
