package enums

type BookStatus int

const (
	Available BookStatus = iota + 1
	ReadyToTake
	Borrowed
	Expired
)

func (b BookStatus) ToString() string {
	switch b {
	case Available:
		return "Available"
	case ReadyToTake:
		return "Ready to Take"
	case Borrowed:
		return "Borrowed"
	case Expired:
		return "Expired"
	}
	return "Unknown"
}
