package serviceUser

import (
	apperrors "auth/internal/errors"
	serviceUserModel "auth/internal/service/user/model"
	"context"
)

func (s *UserServ) RefreshToken(ctx context.Context, token serviceUserModel.TokenPair) (*serviceUserModel.TokenPair, error) {
	err := token.Validate()
	if err != nil {
		return nil, err
	}

	user, err := s.repo.GetSession(ctx, token.RefreshToken)
	if err != nil {
		return nil, err
	}

	if user.Session != token.RefreshToken {
		return nil, apperrors.ErrWrongRefreshToken
	}

	newTokenPair, err := generateTokePair(user, s.JWTSecret())
	if err != nil {
		return nil, err
	}

	return newTokenPair, nil
}
