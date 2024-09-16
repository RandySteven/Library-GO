package responses

import "time"

type (
	UserRegisterResponse struct {
		ID        uint64    `json:"id"`
		Token     string    `json:"token"`
		CreatedAt time.Time `json:"created_at"`
	}

	UserLoginResponse struct {
		Token string `json:"token"`
	}
)
