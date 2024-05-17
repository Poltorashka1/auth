package converter

import (
	repoUserModel "auth/internal/repository/user/model"
	serviceUserModel "auth/internal/service/user/model"
	"strings"
)

func RepoToUsers(users *repoUserModel.Users) *serviceUserModel.Users {
	u := &serviceUserModel.Users{}

	for _, us := range users.Users {
		u.Users = append(u.Users, *RepoToUser(&us))
	}

	return u
}

func RepoToUser(user *repoUserModel.User) *serviceUserModel.User {
	roles := strings.Split(user.Roles, ", ")

	return &serviceUserModel.User{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Active:   user.Active,
		Roles:    roles,
	}
}
