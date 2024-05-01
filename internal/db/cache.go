package db

import "context"

type Cache interface {
	Get(ctx context.Context, key string) (string, error)
	Add(ctx context.Context, key string, value string) error

	Close() error
}
