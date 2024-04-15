package repoUser

import (
	"auth/internal/client/db"
	"context"
)

func (r *UserRepos) ActivateUser(ctx context.Context, email string) error {

	q := db.NewQuery(
		"ActivateUser",
		"update users set active = true where email = $1",
		[]interface{}{email},
	)

	_, err := r.db.ExecContext(ctx, q)
	if err != nil {
		return err
	}
	return nil
}
