package repoUser

import (
	"auth/internal/client/db"
	serviceUserModel "auth/internal/service/user/model"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	"strings"
)

func (r *UserRepos) SignUp(ctx context.Context, user serviceUserModel.SignUpUser, token string) (int64, error) {
	var id int64
	q := db.NewQuery(
		"SignUp",
		"INSERT INTO users (name, email, password, token) VALUES ($1, $2, $3, $4) RETURNING id",
		[]interface{}{user.Name, user.Email, user.Password, token},
	)

	err := r.db.QueryRowContext(ctx, q).Scan(&id)
	if err != nil {
		var pgErr *pgconn.PgError
		switch {
		case errors.As(err, &pgErr):
			if pgErr.Code == "23505" {
				return 0, fmt.Errorf("user with this %s already exists", strings.Split(pgErr.ConstraintName, "_")[1])
			}
		}
		return 0, err
	}

	return id, nil
}

// func (r *repos) SignUp(ctx context.Context, user serviceUserModel.SignUpUser) (int64, error) {
//	// open transaction
//	tx, err := r.db.DB().StartTransaction()
//	if err != nil {
//		return 0, err
//	}
//
//	// transaction into context
//	var ctxValue = context.WithValue(ctx, stx, tx)
//	var id int64
//	var matchField string
//
//	// first request
//	q := db.NewQuery(
//		"SignUp",
//		`
//			SELECT
//				CASE
//					WHEN email = $1 AND name = $2 THEN 'name and email'
//				    WHEN name = $2 THEN 'name'
//					WHEN email = $1 THEN 'email'
//				END AS match_field
//			FROM users
//			WHERE email = $1 OR name = $2;
//				`,
//		[]interface{}{user.Email, user.Name},
//	)
//
//	err = r.db.DB().QueryRowContext(ctxValue, q).Scan(&matchField)
//	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
//		_ = tx.Rollback(ctx)
//		return 0, err
//	}
//
//	if matchField != "" {
//		_ = tx.Rollback(ctx)
//		switch matchField {
//		case "name and email":
//			return 0, apperrors.ErrUserAlreadyExists
//		case "name":
//			return 0, apperrors.ErrUserNameAlreadyExists
//		case "email":
//			return 0, apperrors.ErrUserEmailAlreadyExists
//		}
//	}
//
//	// second request
//	passwd, err := user.Password()
//	if err != nil {
//		_ = tx.Rollback(ctx)
//		return 0, err
//	}
//
//	q = db.NewQuery(
//		"SignUp",
//		"INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id",
//		[]interface{}{user.Name, user.Email, passwd},
//	)
//
//	err = r.db.DB().QueryRowContext(ctxValue, q).Scan(&id)
//	if err != nil {
//		_ = tx.Rollback(ctx)
//		var pgErr *pgconn.PgError
//		switch {
//		case errors.As(err, &pgErr):
//			if pgErr.Code == "23505" {
//				return 0, fmt.Errorf("user with this %s already exists", strings.Split(pgErr.ConstraintName, "_")[1])
//			}
//		}
//		return 0, err
//	}
//
//	err = tx.Commit(ctx)
//	if err != nil {
//		return 0, err
//	}
//
//	return id, nil
//}
