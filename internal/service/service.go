package service

import (
	"auth/internal/db"
	"auth/internal/logger"
	"auth/internal/repository"
	serviceUser "auth/internal/service/user"
	"auth/internal/smtp"
)

type Service interface {
	serviceUser.UserService
}

// container of all services tables and methods to work with it
type service struct {
	serviceUser.UserService
}

func New(repo repository.Repository, tx db.Transaction, smtp smtp.SMTP, log logger.Logger) Service {
	return &service{
		serviceUser.New(repo, tx, smtp, log),
	}
}
