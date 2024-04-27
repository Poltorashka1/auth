package repoUser

import (
	"auth/internal/db"
	"context"
)

func (r *UserRepos) DeleteVerifyToken(ctx context.Context, email string) error {
	q := db.NewQuery(
		"DeleteVerifyToken",
		`UPDATE "Users" SET email_verify_token = NULL WHERE email = $1`,
		[]interface{}{email},
	)

	_, err := r.db.ExecContext(ctx, q)
	if err != nil {
		return err
	}
	return nil
}
