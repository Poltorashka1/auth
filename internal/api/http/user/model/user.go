package apiUserModel

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type SignUpUser struct {
	Username string `json:"username"`
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
	Username  string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	Session   string    `json:"session,omitempty"`
}

type AuthTokenPair struct {
	RefreshToken string
	AccessToken  string
}

type TokenData struct {
	UserName string `json:"username"`
	UserRole string `json:"role"`
	Expires  int    `json:"exp"`
}

type CheckUserRoleData struct {
	Username string `json:"username"`
	UserRole string `json:"role"`
	Route    string `json:"route"`
}

func (t *TokenData) FromClaims(claims jwt.MapClaims) {
	if username, ok := claims["username"].(string); ok {
		t.UserName = username
	}

	if role, ok := claims["role"].(string); ok {
		t.UserRole = role
	}
}
