package serviceUser

import (
	apperrors "auth/internal/errors"
	serviceUserModel "auth/internal/service/user/model"
	"auth/internal/utils"
	"context"
	"strings"
	"time"
)

// todo signin with username

func (s *UserServ) SignIn(ctx context.Context, signIn serviceUserModel.SignInUser) (*serviceUserModel.AuthTokenPair, error) {
	err := signIn.Validate()
	if err != nil {
		return nil, err
	}

	user, err := s.repo.GetUserByEmail(ctx, signIn.Email)
	if err != nil {
		return nil, err
	}

	if !user.Active {
		return nil, apperrors.ErrUserNotActivated
	}

	err = user.CheckPassword(signIn.Password)
	if err != nil {
		return nil, apperrors.ErrWrongPassword
	}

	tokenPair, err := utils.TokenPair(createTokenData(user), s.JWTSecret())
	if err != nil {
		return nil, err
	}

	err = s.repo.CreateSession(ctx, createSession(serviceUserModel.Session{
		UserID:       user.ID,
		Username:     user.Username,
		RefreshToken: tokenPair.RefreshToken,
	}))
	if err != nil {
		return nil, err
	}

	return tokenPair, nil
}

func createTokenData(user *serviceUserModel.User) *serviceUserModel.TokenData {
	return &serviceUserModel.TokenData{
		Username: user.Username,
		UserRole: strings.Join(user.Roles, ", "),
	}
}

func createSession(session serviceUserModel.Session) serviceUserModel.CreateSession {
	return serviceUserModel.CreateSession{
		UserID:       session.UserID,
		Username:     session.Username,
		RefreshToken: session.RefreshToken,

		// todo IMPORTANT time not work because it dont need now
		Expires: time.Now(),
	}
}
