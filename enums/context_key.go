package enums

type ContextKey string

const (
	UserID      ContextKey = `user_id`
	RoleID                 = `role_id`
	RequestID              = `request_id`
	Env                    = `env`
	ClientIP               = `client_ip`
	RandomState            = `randomstateoauth2`
)

func (c ContextKey) ToString() string {
	return string(c)
}
