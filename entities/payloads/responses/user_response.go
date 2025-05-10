package responses

import "time"

type (
	UserRegisterResponse struct {
		ID        string    `json:"id"`
		Token     string    `json:"token"`
		CreatedAt time.Time `json:"created_at"`
	}

	Oauth2Response struct {
		ID            string `json:"id"`
		Email         string `json:"email"`
		VerifiedEmail bool   `json:"verified_email"`
		Name          string `json:"name"`
		GivenName     string `json:"given_name"`
		FamilyName    string `json:"family_name"`
		Picture       string `json:"picture"`
		Locale        string `json:"locale"`
	}

	UserLoginResponse struct {
		Token string `json:"token"`
	}

	UserListResponse struct {
		ID    uint64 `json:"id"`
		Name  string `json:"name"`
		Email string `json:"consumers"`
	}

	UserDetailResponse struct {
		ID             uint64    `json:"id"`
		Name           string    `json:"name"`
		Address        string    `json:"address"`
		PhoneNumber    string    `json:"phone_number"`
		ProfilePicture string    `json:"profile_picture"`
		DoB            time.Time `json:"date_of_birth"`
	}

	LoginUserResponse struct {
		ID             uint64    `json:"id"`
		RoleID         uint64    `json:"role_id"`
		Name           string    `json:"name"`
		Address        string    `json:"address"`
		Email          string    `json:"consumers"`
		PhoneNumber    string    `json:"phone_number"`
		ProfilePicture string    `json:"profile_picture"`
		Password       string    `json:"password"`
		DoB            time.Time `json:"date_of_birth"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
	}
)
