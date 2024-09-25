package enums

type BookStatus int

const (
	Available BookStatus = iota + 1
	ReadyToTake
	Borrowed
)

func (b BookStatus) ToString() string {
	switch b {
	case Available:
		return "Available"
	case ReadyToTake:
		return "Ready to Take"
	case Borrowed:
		return "Borrowed"
	}
	return "Unknown"
}
