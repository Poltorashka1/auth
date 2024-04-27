package serviceUser

import (
	apperrors "auth/internal/errors"
	serviceUserModel "auth/internal/service/user/model"
	"context"
)

// todo обработать все ошибки

func (s *UserServ) EmailVerify(ctx context.Context, verify serviceUserModel.EmailVerify) error {
	err := verify.Validate()
	if err != nil {
		return err
	}

	verifyToken, err := s.repo.CheckUserVerify(ctx, verify.Email)
	if err != nil {
		return err
	}

	if verifyToken != verify.Token {
		return apperrors.ErrWrongVerifyToken
	}

	err = s.repo.ActivateUser(ctx, verify.Email)
	if err != nil {
		return err
	}

	err = s.repo.DeleteVerifyToken(ctx, verify.Email)
	if err != nil {
		return err
	}

	return nil
}
