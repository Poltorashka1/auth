package serviceUser

import (
	apperrors "auth/internal/errors"
	serviceUserModel "auth/internal/service/user/model"
	"context"
)

func (s *UserServ) EmailVerify(ctx context.Context, verify serviceUserModel.EmailVerify) error {
	err := verify.Validate()
	if err != nil {
		return err
	}

	token, err := s.repo.CheckUserVerifyByEmail(ctx, verify.Email)
	if err != nil {
		return err
	}

	if token != verify.Token {
		return apperrors.ErrWrongVerifyToken
	}

	err = s.repo.ActivateUser(ctx, verify.Email)
	if err != nil {
		return err
	}

	return nil
}
