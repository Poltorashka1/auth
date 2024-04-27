package repoUser

import (
	"auth/internal/db"
	"context"
)

func (r *UserRepos) SignOut(ctx context.Context, refreshToken string) error {
	q := db.NewQuery(
		"SignOut",
		`update "Session" set session_token = null where session_token = $1`,
		[]interface{}{refreshToken},
	)

	// todo if сессии нету?
	_, err := r.db.ExecContext(ctx, q)
	if err != nil {
		return err
	}
	return nil
}
