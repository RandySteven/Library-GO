package requests

type (
	UserRegisterRequest struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Password  string `json:"password"`
		UserName  string `json:"username"`
	}

	UserLoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)
