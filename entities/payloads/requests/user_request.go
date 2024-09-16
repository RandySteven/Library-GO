package requests

type (
	UserRegisterRequest struct {
		FirstName   string `json:"first_name" validate:"required"`
		LastName    string `json:"last_name" validate:"required"`
		PhoneNumber string `json:"phone_number" validate:"required"`
		Email       string `json:"email" validate:"required,email"`
		Password    string `json:"password" validate:"required"`
		Day         string `json:"day" validate:"required"`
		Month       string `json:"month" validate:"required"`
		Year        string `json:"year" validate:"required"`
	}

	UserLoginRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	VerifyUserRequest struct {
		Token string `json:"token" validate:"required"`
	}
)
