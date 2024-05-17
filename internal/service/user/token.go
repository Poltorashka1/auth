package serviceUser

import (
	serviceUserModel "auth/internal/service/user/model"
	"auth/internal/utils"
	"context"
)

// todo возможно стоит передавать id пользователя

// AccessToken - errors : apperrors.ErrWrongRefreshToken,
func (s *UserServ) AccessToken(ctx context.Context, refreshToken serviceUserModel.RefreshToken) (*serviceUserModel.AuthTokenPair, error) {
	err := refreshToken.Validate()
	if err != nil {
		return nil, err
	}

	// todo delete но не факт если не найху решение
	session, err := s.repo.Session(ctx, refreshToken.RefreshToken)
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

	ses := serviceUserModel.CreateSession{
		UserID:       session.UserID,
		Username:     session.Username,
		RefreshToken: newTokenPair.RefreshToken,
		// todo
		//Expires:
	}

	err = s.repo.CreateSession(ctx, ses)
	if err != nil {
		return nil, err
	}

	return newTokenPair, nil
}
