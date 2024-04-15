package serviceUser

import (
	apperrors "auth/internal/errors"
	serviceUserModel "auth/internal/service/user/model"
	"context"
)

func (s *UserServ) SignIn(ctx context.Context, user serviceUserModel.SignInUser) (string, error) {
	err := user.Validate()
	if err != nil {
		return "", err
	}

	// todo mb pgx error lovit tut
	// user exist
	dbUser, err := s.repo.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return "", err
	}

	if !dbUser.Active {
		return "", apperrors.ErrUserNotActivated
	}

	err = user.CheckPassword(dbUser.Password)
	if err != nil {
		return "", apperrors.ErrWrongPassword
	}

	token, err := dbUser.GenerateJwtToken(s.JWTSecret())
	if err != nil {
		return "", err
	}

	return token, nil
}
