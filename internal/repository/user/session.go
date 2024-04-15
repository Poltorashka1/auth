package repoUser

import (
	"auth/internal/client/db"
	serviceModel "auth/internal/service/user/model"
	"context"
)

func (r *UserRepos) CreateSession(ctx context.Context, id int, refreshToken string) error {
	q := db.NewQuery(
		"CreateSession",
		"UPDATE users SET session = $2 WHERE id = $1",
		[]interface{}{id, refreshToken},
	)

	_, err := r.db.ExecContext(ctx, q)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepos) GetSession(ctx context.Context, refreshToken string) (*serviceModel.User, error) {
	q := db.NewQuery(
		"GetSession",
		"SELECT session FROM users WHERE session = $1",
		[]interface{}{refreshToken},
	)

	var user = new(serviceModel.User)
	err := r.db.ScanOneContext(ctx, user, q)
	if err != nil {
		return nil, err
	}

	return user, nil
}
