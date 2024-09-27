package enums

type RouterPrefix string

const (
	DevPrefix        RouterPrefix = "/dev"
	OnboardingPrefix RouterPrefix = "/onboarding"
	UserPrefix       RouterPrefix = "/users"
	BookPrefix       RouterPrefix = "/books"
	GenrePrefix      RouterPrefix = "/genres"
	BagPrefix        RouterPrefix = "/bags"
	StoryPrefix      RouterPrefix = "/stories"
	BorrowPrefix     RouterPrefix = "/borrows"
)

func (r RouterPrefix) ToString() string {
	return string(r)
}
