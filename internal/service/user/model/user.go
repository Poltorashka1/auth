package serviceUserModel

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type TokenPair struct {
	RefreshToken string
	AccessToken  string
}

type GetUserByNameOrID struct {
	Param string
	Field string
}

type EmailVerify struct {
	Token string
	Email string
}

type SignUpUser struct {
	Name     string
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

func (s *SignInUser) CheckPassword(hashPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(s.Password))
}

type User struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Role      int       `db:"role"`
	Active    bool      `db:"active"`
	CreatedAt time.Time `db:"created_at"`
	Session   string    `db:"session"`
}

func (u *User) Validate() error {
	return nil
}
