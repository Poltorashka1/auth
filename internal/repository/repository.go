package repository

import (
	"auth/internal/db"
	"auth/internal/logger"
	repositoryUser "auth/internal/repository/user"
)

type Repository interface {
	repositoryUser.UserRepository
}

// container of all repositories tables and methods to work with it
type repository struct {
	repositoryUser.UserRepository
}

func New(db db.DB, log logger.Logger) Repository {
	return &repository{
		repositoryUser.New(db, log),
	}
}
