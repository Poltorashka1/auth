package repoUser

import (
	"auth/internal/db"
	apperrors "auth/internal/errors"
	serviceUserModel "auth/internal/service/user/model"
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
)

func (r *UserRepos) CheckUserByNameAndEmail(ctx context.Context, user serviceUserModel.SignUpUser) error {
	var matchField string

	q := db.NewQuery(
		"GetUserByNameAndEmail",
		`
			SELECT
				CASE
					WHEN email = $1 AND username = $2 THEN 'username and email'
				    WHEN username = $2 THEN 'username'
					WHEN email = $1 THEN 'email'
				END AS match_field
			FROM "Users"
			WHERE email = $1 OR username = $2;
				`,
		[]interface{}{user.Email, user.Username},
	)

	err := r.db.QueryRowContext(ctx, q).Scan(&matchField)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return err
	}

	if matchField != "" {
		switch matchField {
		case "username and email":
			return apperrors.ErrUserAlreadyExists
		case "username":
			return apperrors.ErrUserNameAlreadyExists
		case "email":
			return apperrors.ErrUserEmailAlreadyExists
		}
	}
	return nil
}

func (r *UserRepos) CheckUserVerify(ctx context.Context, email string) (verifyToken string, err error) {
	q := db.Query{
		Name:     "CheckUserVerifyByEmail",
		QueryRaw: `select active, email_verify_token from "Users" where email = $1`,
		Args:     []interface{}{email},
	}
	var active bool
	var token string
	err = r.db.QueryRowContext(ctx, q).Scan(&active, &token)
	if err != nil {
		return "", err
	}

	if active {
		return "", apperrors.ErrUserAlreadyActive
	}

	//if token == "" {
	//	return "", apperrors.ErrUserNotVerified
	//}
	return token, nil
}
