package serviceUser

import (
	serviceUserModel "auth/internal/service/user/model"
	"context"
	"log"
)

func (s *UserServ) SignUp(ctx context.Context, user serviceUserModel.SignUpUser) (int64, error) {
	err := user.Validate()
	if err != nil {
		return 0, err
	}

	txCtx, err := s.tx.StartTransaction(ctx)
	if err != nil {
		return 0, err
	}

	defer func() {
		if err != nil {
			err := s.tx.Rollback(txCtx)
			if err != nil {
				log.Println(err)
			}
		}
	}()

	err = s.repo.CheckUserByNameAndEmail(txCtx, user)
	if err != nil {
		return 0, err
	}

	err = user.HashPassword()
	if err != nil {
		return 0, err
	}

	// generate verification token
	token, err := s.VerifyToken()
	if err != nil {
		return 0, err
	}

	id, err := s.repo.SignUp(txCtx, user, token)
	if err != nil {
		return 0, err
	}

	err = s.smtp.SendEmail(user.Email, token, user.Name)
	if err != nil {
		return 0, err
	}

	err = s.tx.Commit(txCtx)
	if err != nil {
		return 0, err
	}

	return id, nil
}
