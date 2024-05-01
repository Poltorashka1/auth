package repoUser

import (
	"auth/internal/db"
	"context"
)

func (r *UserRepos) AddRole(ctx context.Context, userID int64, roleID int) error {
	q := db.NewQuery(
		"AddRole",
		`INSERT INTO "UserRoles" (user_id, role_id) VALUES ($1, $2)`,
		[]interface{}{userID, roleID})

	_, err := r.db.ExecContext(ctx, q)
	if err != nil {
		return err
	}
	return nil
}
