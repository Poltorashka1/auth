package serviceUser

import (
	serviceUserModel "auth/internal/service/user/model"
	"context"
	"crypto/rand"
	"encoding/base64"
	"log"
)

// SignUp create new user and send verification email.
// errors: apperrors.ValidationErrors, pgx.ErrNoRows, apperrors.ExistsError, apperrors.ErrSMTPSendMessage
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
	token, err := generateVerifyToken()
	if err != nil {
		return 0, err
	}

	id, err := s.repo.CreateUser(txCtx, user, token)
	if err != nil {
		return 0, err
	}

	// todo edit 2 на что то более понятное
	err = s.repo.AddRole(txCtx, id, 2)
	if err != nil {
		return 0, err
	}

	err = s.smtp.SendEmail(user.Email, token, user.Username)
	if err != nil {
		return 0, err
	}

	err = s.tx.Commit(txCtx)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func generateVerifyToken() (string, error) {
	bt := make([]byte, 32)
	_, err := rand.Read(bt)
	if err != nil {
		return "", err
	}

	token := base64.URLEncoding.EncodeToString(bt)
	return token, err
}
