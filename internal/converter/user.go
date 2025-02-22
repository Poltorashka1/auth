package converter

import (
	apiUSerModel "auth/internal/api/http/user/model"
	serviceUserModel "auth/internal/service/user/model"
	"auth/pkg"
)

func FromPkgToUserService(user *pkg.SignUpRequest) serviceUserModel.SignUpUser {
	return serviceUserModel.SignUpUser{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}

func FromHTTPToUserService(user *apiUSerModel.SignUpUser) serviceUserModel.SignUpUser {
	return serviceUserModel.SignUpUser{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}

func FromAPIToGetUser(field, param string) serviceUserModel.GetUserByNameOrID {
	return serviceUserModel.GetUserByNameOrID{
		Field: field,
		Param: param,
	}
}

func FromApiToEmailVerify(req apiUSerModel.EmailVerify) serviceUserModel.EmailVerify {
	return serviceUserModel.EmailVerify{
		Token: req.Token,
		Email: req.Email,
	}
}

func FromApiToSignIn(req apiUSerModel.SignInUser) serviceUserModel.SignInUser {
	return serviceUserModel.SignInUser{
		Email:    req.Email,
		Password: req.Password,
	}
}
