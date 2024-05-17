package serviceUser

import (
	serviceUserModel "auth/internal/service/user/model"
	"context"
)

func (s *UserServ) UserByID(ctx context.Context, id serviceUserModel.UserByID) (*serviceUserModel.User, error) {
	err := id.Validate()
	if err != nil {
		return nil, err
	}

	user, err := s.repo.UserByID(ctx, id.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserServ) UserByName(ctx context.Context, username serviceUserModel.UserByName) (*serviceUserModel.User, error) {
	err := username.Validate()
	if err != nil {
		return nil, err
	}

	user, err := s.repo.UserByName(ctx, username.Name)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserServ) Users(ctx context.Context) (*serviceUserModel.Users, error) {
	users, err := s.repo.Users(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}
