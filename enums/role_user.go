package enums

type RoleUser int

const (
	Admin RoleUser = iota + 1
	Member
	Premium
	Librarian
)
