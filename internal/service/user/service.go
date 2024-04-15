package serviceUser

import (
	"auth/internal/client/db"
	"auth/internal/repository"
	serviceUserModel "auth/internal/service/user/model"
	"auth/internal/smtp"
	"context"
	"log"
	"os"
)

type UserService interface {
	SignUp(ctx context.Context, user serviceUserModel.SignUpUser) (int64, error)
	SignIn(ctx context.Context, user serviceUserModel.SignInUser) (*serviceUserModel.TokenPair, error)
	// SignOut(ctx context.Context, token string) error

	GetUser(ctx context.Context, idOrName serviceUserModel.GetUserByNameOrID) (*serviceUserModel.User, error)
	EmailVerify(ctx context.Context, verify serviceUserModel.EmailVerify) error
	RefreshToken(ctx context.Context, token serviceUserModel.TokenPair) (*serviceUserModel.TokenPair, error)
}

type UserServ struct {
	smtp       smtp.SMTP
	repo       repository.Repository
	tx         db.Transaction
	_jwtSecret string
}

func (s *UserServ) JWTSecret() string {
	return s._jwtSecret
}

func New(userRepo repository.Repository, tx db.Transaction, smtp smtp.SMTP) UserService {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		// todo logging
		log.Fatal("JWT_SECRET is not set")
	}

	// todo logging
	return &UserServ{
		repo:       userRepo,
		tx:         tx,
		smtp:       smtp,
		_jwtSecret: secret,
	}
}
