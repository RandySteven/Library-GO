package _const

type ContextKey string

const (
	RequestID ContextKey = "RequestID"
	UserID    ContextKey = "UserID"
	RoleID    ContextKey = "RoleID"
)
