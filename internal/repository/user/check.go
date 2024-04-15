package repoUser

import (
	"auth/internal/client/db"
	apperrors "auth/internal/errors"
	serviceUserModel "auth/internal/service/user/model"
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
)

// todo mb nado vozvrashat field?

func (r *UserRepos) CheckUserByNameAndEmail(ctx context.Context, user serviceUserModel.SignUpUser) error {
	var matchField string

	q := db.NewQuery(
		"GetUserByNameAndEmail",
		`
			SELECT
				CASE
					WHEN email = $1 AND name = $2 THEN 'name and email'
				    WHEN name = $2 THEN 'name'
					WHEN email = $1 THEN 'email'
				END AS match_field
			FROM users
			WHERE email = $1 OR name = $2;
				`,
		[]interface{}{user.Email, user.Name},
	)

	err := r.db.QueryRowContext(ctx, q).Scan(&matchField)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return err
	}

	if matchField != "" {
		switch matchField {
		case "name and email":
			return apperrors.ErrUserAlreadyExists
		case "name":
			return apperrors.ErrUserNameAlreadyExists
		case "email":
			return apperrors.ErrUserEmailAlreadyExists
		}
	}
	return nil
}

// todo mb nado vozvrashat active?

func (r *UserRepos) CheckUserVerifyByEmail(ctx context.Context, email string) (string, error) {
	q := db.Query{
		Name:     "CheckUserVerifyByEmail",
		QueryRaw: "select active, token from users where email = $1",
		Args:     []interface{}{email},
	}
	var active bool
	var token string
	err := r.db.QueryRowContext(ctx, q).Scan(&active, &token)
	if err != nil {
		return "", err
	}

	if active {
		return "", apperrors.ErrUserAlreadyActive
	}
	return token, nil
}
