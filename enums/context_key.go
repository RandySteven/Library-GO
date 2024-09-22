package enums

type ContextKey string

const (
	UserID    ContextKey = `user_id`
	RoleID               = `role_id`
	RequestID            = `request_id`
)

func (c ContextKey) ToString() string {
	return string(c)
}
