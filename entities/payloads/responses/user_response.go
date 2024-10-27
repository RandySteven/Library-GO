package responses

import "time"

type (
	UserRegisterResponse struct {
		ID        string    `json:"id"`
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
		ID          uint64    `json:"id"`
		Name        string    `json:"name"`
		Address     string    `json:"address"`
		Email       string    `json:"email"`
		PhoneNumber string    `json:"phone_number"`
		Password    string    `json:"password"`
		DoB         time.Time `json:"date_of_birth"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}
)
