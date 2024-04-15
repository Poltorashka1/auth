package repoUser

import (
	"auth/internal/client/db"
	"auth/internal/logger"
	serviceUserModel "auth/internal/service/user/model"
	"context"
)

type UserRepository interface {
	SignUp(ctx context.Context, user serviceUserModel.SignUpUser, token string) (int64, error)
	GetUserByNameOrID(ctx context.Context, param serviceUserModel.GetUserByNameOrID) (*serviceUserModel.User, error)
	CheckUserByNameAndEmail(ctx context.Context, param serviceUserModel.SignUpUser) error
	CheckUserVerifyByEmail(ctx context.Context, email string) (string, error)
	ActivateUser(ctx context.Context, email string) error
	SignIn(ctx context.Context, user serviceUserModel.SignInUser) (string, error)
	GetUserByEmail(ctx context.Context, email string) (*serviceUserModel.User, error)
	CreateSession(ctx context.Context, id int, refreshToken string) error
	// todo поменять
	GetSession(ctx context.Context, refreshToken string) (*serviceUserModel.User, error)
	// SignOut(ctx context.Context, token string) error

	// GetUserByName(ctx context.Context, user serviceUserModel.GetUser) (*serviceUserModel.User, error)
	// GetUserByID(ctx context.Context, user serviceUserModel.GetUser) (*serviceUserModel.User, error)
}

type UserRepos struct {
	// db  db.Client
	db  db.DB
	log logger.Logger
}

func New(db db.DB, log logger.Logger) *UserRepos {
	return &UserRepos{
		db:  db,
		log: log,
	}
}

// func (r *repos) SignIn(ctx context.Context, user serviceUserModel.SignInUser) (string, error) {
//	return "", nil
//}
//
// func (r *repos) SignOut(ctx context.Context, token string) error {
//	return nil
//}
//
