package repoUser

import (
	"auth/internal/db"
	"context"
)

func (r *UserRepos) ActivateUser(ctx context.Context, email string) error {

	q := db.NewQuery(
		"ActivateUser",
		`update "Users" set active = true where email = $1`,
		[]interface{}{email},
	)

	_, err := r.db.ExecContext(ctx, q)
	if err != nil {
		return err
	}
	return nil
}
