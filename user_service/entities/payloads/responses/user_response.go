package responses

import "time"

type (
	LoginToken string

	UserProfile struct {
		FullName    string `json:"full_name"`
		Username    string `json:"username"`
		Email       string `json:"email"`
		Password    string `json:"password"`
		DateOfBirth string `json:"date_of_birth"`
	}

	UserDetail struct {
		Username    string `json:"username"`
		FullName    string `json:"full_name"`
		Email       string `json:"email"`
		DateOfBirth string `json:"date_of_birth"`
	}

	RegisterResponse struct {
		ID          uint64    `json:"id"`
		FullName    string    `json:"full_name"`
		FirstName   string    `json:"first_name"`
		LastName    string    `json:"last_name"`
		Email       string    `json:"email"`
		Username    string    `json:"username"`
		DateOfBirth string    `json:"date_of_birth"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
		CreatedBy   uint64    `json:"created_by"`
		UpdatedBy   uint64    `json:"updated_by"`
	}
)
