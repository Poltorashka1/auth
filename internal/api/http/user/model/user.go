package apiUserModel

import "time"

type SignUpUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type EmailVerify struct {
	Token string
	Email string
}

type SignUpResponse struct {
	ID int64 `json:"id"`
}

type SignInUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	Session   string    `json:"session"`
}

type TokenPair struct {
	RefreshToken string
	AccessToken  string
}
