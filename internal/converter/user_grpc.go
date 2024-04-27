package converter

import (
	serviceUserModel "auth/internal/service/user/model"
	"auth/pkg"
)

func GRPCToSignUp(req *pkg.SignUpRequest) serviceUserModel.SignUpUser {
	return serviceUserModel.SignUpUser{
		Username: req.GetName(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}
}
