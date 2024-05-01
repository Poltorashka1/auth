package apiUser

import (
	"auth/internal/logger"
	"auth/internal/service"
)

type UserHandler struct {
	serv       service.Service
	log        logger.Logger
	_jwtSecret string
}

func New(serv service.Service, log logger.Logger, jwtSecret string) *UserHandler {
	return &UserHandler{
		serv:       serv,
		log:        log,
		_jwtSecret: jwtSecret,
	}
}
