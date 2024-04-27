package serviceUser

import (
	serviceUserModel "auth/internal/service/user/model"
	"context"
)

func (s *UserServ) GetUserByID(ctx context.Context, userID serviceUserModel.GetUserByID) (*serviceUserModel.User, error) {
	err := userID.Validate()
	if err != nil {
		return nil, err
	}

	user, err := s.repo.GetUserByID(ctx, userID.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserServ) GetUserByName(ctx context.Context, userName serviceUserModel.GetUserByName) (*serviceUserModel.User, error) {
	err := userName.Validate()
	if err != nil {
		return nil, err
	}

	user, err := s.repo.GetUserByName(ctx, userName.Name)
	if err != nil {
		return nil, err
	}

	return user, nil
}
