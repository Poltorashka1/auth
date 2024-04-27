package repoUser

import (
	"auth/internal/db"
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
		`INSERT INTO "Users" (username, email, password, email_verify_token) VALUES ($1, $2, $3, $4) RETURNING id`,
		[]interface{}{user.Username, user.Email, user.Password, token},
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
