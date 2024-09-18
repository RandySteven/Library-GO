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

	UserListResponse struct {
		ID    uint64 `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	UserDetailResponse struct {
	}

	LoginUserResponse struct {
	}
)
