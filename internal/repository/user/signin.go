package repoUser

import (
	serviceUserModel "auth/internal/service/user/model"
	"context"
)

func (r *UserRepos) SignIn(ctx context.Context, user serviceUserModel.SignInUser) (string, error) {
	return "", nil
}
