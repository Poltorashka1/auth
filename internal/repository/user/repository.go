package repoUser

import (
	"auth/internal/db"
	"auth/internal/logger"
	serviceUserModel "auth/internal/service/user/model"
	"context"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user serviceUserModel.SignUpUser, token string) (int64, error)

	SignIn(ctx context.Context, user serviceUserModel.SignInUser) (string, error)
	SignOut(ctx context.Context, refreshToken string) error

	GetUserByID(ctx context.Context, id string) (*serviceUserModel.User, error)
	GetUserByName(ctx context.Context, username string) (*serviceUserModel.User, error)
	GetUserByEmail(ctx context.Context, email string) (*serviceUserModel.User, error)
	GetUserRoles(ctx context.Context, userID int) (string, error)

	CheckUserByNameAndEmail(ctx context.Context, param serviceUserModel.SignUpUser) error
	CheckUserVerify(ctx context.Context, email string) (string, error)

	ActivateUser(ctx context.Context, email string) error
	AddRole(ctx context.Context, userID int64, roleID int) error

	CreateSession(ctx context.Context, session serviceUserModel.CreateSession) error
	GetSession(ctx context.Context, refreshToken string) (*serviceUserModel.Session, error)
	GetRouteRole(ctx context.Context, data serviceUserModel.CheckUserRoleData) ([]string, error)

	DeleteVerifyToken(ctx context.Context, email string) error
}

type UserRepos struct {
	// db  db.Client
	db    db.DB
	cache db.Cache
	log   logger.Logger
}

func New(db db.DB, cache db.Cache, log logger.Logger) UserRepository {
	return &UserRepos{
		db:    db,
		cache: cache,
		log:   log,
	}
}
