package enums

type RouterPrefix string

const (
	OnboardingPrefix RouterPrefix = "/onboarding"
	UserPrefix       RouterPrefix = "/users"
	BookPrefix       RouterPrefix = "/books"
)

func (r RouterPrefix) ToString() string {
	return string(r)
}
