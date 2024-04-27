package converter

import (
	apiUserModel "auth/internal/api/http/user/model"
	serviceUserModel "auth/internal/service/user/model"
	"strings"
)

func HTTPToSignUpUser(user *apiUserModel.SignUpUser) serviceUserModel.SignUpUser {
	return serviceUserModel.SignUpUser{
		Username: strings.ToLower(user.Username),
		Email:    user.Email,
		Password: user.Password,
	}
}

func HTTPToEmailVerify(req apiUserModel.EmailVerify) serviceUserModel.EmailVerify {
	return serviceUserModel.EmailVerify{
		Token: req.Token,
		Email: req.Email,
	}
}

func HTTPToSignIn(req apiUserModel.SignInUser) serviceUserModel.SignInUser {
	return serviceUserModel.SignInUser{
		Email:    req.Email,
		Password: req.Password,
	}
}

func HTTPToRefreshToken(inp *apiUserModel.AuthTokenPair) serviceUserModel.AuthTokenPair {
	return serviceUserModel.AuthTokenPair{
		RefreshToken: inp.RefreshToken,
	}
}

func ToHTTPRefreshToken(inp *serviceUserModel.AuthTokenPair) apiUserModel.AuthTokenPair {
	return apiUserModel.AuthTokenPair{
		AccessToken:  inp.AccessToken,
		RefreshToken: inp.RefreshToken,
	}
}

func ToHTTPUser(user *serviceUserModel.User) *apiUserModel.User {
	return &apiUserModel.User{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Role:      strings.Join(user.Roles, ", "),
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
	}
}

func HTTPToGetUserByName(name string) serviceUserModel.GetUserByName {
	return serviceUserModel.GetUserByName{Name: name}
}
func HTTPToGetUserByID(id string) serviceUserModel.GetUserByID {
	return serviceUserModel.GetUserByID{ID: id}
}

func HTTPToSignOut(token string) serviceUserModel.SignOut {
	return serviceUserModel.SignOut{RefreshToken: token}
}

func ToCheckUserRole(data *apiUserModel.CheckUserRoleData) serviceUserModel.CheckUserRoleData {
	return serviceUserModel.CheckUserRoleData{
		Username: data.Username,
		UserRole: data.UserRole,
		Route:    data.Route,
	}
}

func ToHTTPTokenPair(tokenPair *serviceUserModel.AuthTokenPair) *apiUserModel.AuthTokenPair {
	return &apiUserModel.AuthTokenPair{AccessToken: tokenPair.AccessToken, RefreshToken: tokenPair.RefreshToken}
}
