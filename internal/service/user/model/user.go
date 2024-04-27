package serviceUserModel

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthTokenPair struct {
	RefreshToken string
	AccessToken  string
}

type GetUserByID struct {
	ID string
}

type GetUserByName struct {
	Name string
}

type EmailVerify struct {
	Token string
	Email string
}

type SignUpUser struct {
	Username string
	Email    string
	Password string
}

func (s *SignUpUser) HashPassword() error {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(s.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	// todo test this byte
	s.Password = string(hashPass)
	return nil
}

type SignInUser struct {
	Email    string
	Password string
}

type TokenData struct {
	Username string `db:"username"`
	UserRole string `db:"role"`
	Expires  int    `db:"exp"`
}

type CheckUserRoleData struct {
	Username string `db:"username"`
	UserRole string `db:"role"`
	Route    string `db:"route"`
}

// User - user from data base
type User struct {
	// from table Users
	ID        int       `db:"id"`
	Username  string    `db:"username"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Active    bool      `db:"active"`
	CreatedAt time.Time `db:"created_at"`
	// from table UserRoles
	Roles []string `db:"roles"`
}

func (u *User) CheckPassword(passwordToCheck string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(passwordToCheck))
}

type Session struct {
	UserID       int    `db:"user_id"`
	Username     string `db:"name"`
	RefreshToken string `db:"session_token"`
}

func (u *User) Validate() error {
	// todo made this func
	return nil
}

type SignOut struct {
	RefreshToken string
}

type CreateSession struct {
	UserID       int
	Username     string
	RefreshToken string
	Expires      time.Time
}
