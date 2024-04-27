package repoUser

import (
	"auth/internal/db"
	apperrors "auth/internal/errors"
	serviceModel "auth/internal/service/user/model"
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
)

func (r *UserRepos) CreateSession(ctx context.Context, session serviceModel.CreateSession) error {
	q := db.NewQuery(
		"CreateSession",
		`
					insert into "Session" (user_id, username, session_token, expiry)
					values ($1, $2, $3, $4) 
					on conflict (user_id) 
					    DO UPDATE
							Set session_token = $3, expiry = $4`,
		[]interface{}{session.UserID, session.Username, session.RefreshToken, session.Expires},
	)

	_, err := r.db.ExecContext(ctx, q)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepos) GetSession(ctx context.Context, refreshToken string) (*serviceModel.Session, error) {
	q := db.NewQuery(
		"GetSession",
		`
					select * from "Session"
					where session_token = $1
					`,
		[]interface{}{refreshToken},
	)

	// todo UserTokenModel
	var session = new(serviceModel.Session)
	err := r.db.ScanOneContext(ctx, session, q)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperrors.ErrWrongRefreshToken
		}
		return nil, err
	}

	return session, nil
}
