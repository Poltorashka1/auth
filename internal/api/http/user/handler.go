package apiUser

import (
	"auth/internal/logger"
	"auth/internal/service"
)

type UserHandler struct {
	serv service.Service
	log  logger.Logger
}

func New(serv service.Service, log logger.Logger) *UserHandler {
	return &UserHandler{
		serv: serv,
		log:  log,
	}
}
