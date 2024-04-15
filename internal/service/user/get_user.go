package serviceUser

import (
	serviceUserModel "auth/internal/service/user/model"
	"context"
)

func (s *UserServ) GetUser(ctx context.Context, idOrName serviceUserModel.GetUserByNameOrID) (*serviceUserModel.User, error) {
	err := idOrName.Validate()
	if err != nil {
		return nil, err
	}

	user, err := s.repo.GetUserByNameOrID(ctx, idOrName)
	if err != nil {
		return nil, err
	}

	return user, nil
}
