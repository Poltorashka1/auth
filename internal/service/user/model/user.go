package serviceUserModel

import (
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"time"
)

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
}

func (u *User) Validate() error {
	return nil
}

func (u *User) GenerateJwtToken(secret string) (string, error) {
	payload := jwt.MapClaims{
		"user_id": u.ID,
		"role":    u.Role,
		"exp":     time.Now().Add(time.Minute * 15).Unix(),
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	token, err := t.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return token, nil
}
