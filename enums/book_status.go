package enums

type BookStatus int

const (
	Available BookStatus = iota + 1
	AtBag
	ReadyToTake
	Borrowed
	Expired
)

func (b BookStatus) ToString() string {
	switch b {
	case Available:
		return "Available"
	case AtBag:
		return "At Bag"
	case ReadyToTake:
		return "Ready to Take"
	case Borrowed:
		return "Borrowed"
	case Expired:
		return "Expired"
	}
	return "Unknown"
}
