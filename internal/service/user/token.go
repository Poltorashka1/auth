package serviceUser

import (
	serviceUserModel "auth/internal/service/user/model"
	"auth/internal/utils"
	"context"
)

func (s *UserServ) GetAccessToken(ctx context.Context, token serviceUserModel.AuthTokenPair) (*serviceUserModel.AuthTokenPair, error) {
	err := token.Validate()
	if err != nil {
		return nil, err
	}

	session, err := s.repo.GetSession(ctx, token.RefreshToken)
	if err != nil {
		return nil, err
	}

	userRoles, err := s.repo.GetUserRoles(ctx, session.UserID)
	if err != nil {
		return nil, err
	}

	tokenData := &serviceUserModel.TokenData{
		Username: session.Username,
		UserRole: userRoles,
	}

	newTokenPair, err := utils.TokenPair(tokenData, s.JWTSecret())
	if err != nil {
		return nil, err
	}

	err = s.repo.CreateSession(ctx, createSession(*session))
	if err != nil {
		return nil, err
	}
	return newTokenPair, nil
}
