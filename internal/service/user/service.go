package serviceUser

import (
	"auth/internal/db"
	"auth/internal/repository"
	serviceUserModel "auth/internal/service/user/model"
	"auth/internal/smtp"
	"context"
	"log"
	"os"
)

type UserService interface {
	SignUp(ctx context.Context, user serviceUserModel.SignUpUser) (int64, error)
	SignIn(ctx context.Context, user serviceUserModel.SignInUser) (*serviceUserModel.AuthTokenPair, error)
	SignOut(ctx context.Context, token serviceUserModel.SignOut) error

	GetUserByID(ctx context.Context, userID serviceUserModel.GetUserByID) (*serviceUserModel.User, error)
	GetUserByName(ctx context.Context, userName serviceUserModel.GetUserByName) (*serviceUserModel.User, error)

	EmailVerify(ctx context.Context, verify serviceUserModel.EmailVerify) error

	GetAccessToken(ctx context.Context, token serviceUserModel.AuthTokenPair) (*serviceUserModel.AuthTokenPair, error)
	CheckUserRole(context context.Context, data serviceUserModel.CheckUserRoleData) error
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
