package repoUser

import (
	"auth/internal/db"
	apperrors "auth/internal/errors"
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
)

func (r *UserRepos) RouteRole(ctx context.Context, route string) (string, error) {
	const op = "RouteRole"

	q := db.CacheQuery{
		Name: op,
		Key:  route,
	}

	result, err := r.cache.Get(ctx, q)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", apperrors.ErrRouteNotFound
		}
		return "", err
	}

	return result, nil
}
