package serviceUser

import (
	serviceUserModel "auth/internal/service/user/model"
	"context"
)

func (s *UserServ) SignOut(ctx context.Context, token serviceUserModel.SignOut) error {
	err := token.Validate()
	if err != nil {
		return err
	}

	err = s.repo.SignOut(ctx, token.RefreshToken)
	if err != nil {
		return err
	}

	return nil
}
