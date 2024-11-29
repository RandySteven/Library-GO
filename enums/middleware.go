package enums

type Middleware int

const (
	AuthenticationMiddleware Middleware = iota + 1
	RateLimiterMiddleware
)
