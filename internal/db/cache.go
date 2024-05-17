package db

import (
	"context"
)

type Cache interface {
	Get(ctx context.Context, q CacheQuery) (string, error)
	Add(ctx context.Context, q CacheQuery) error
	Close() error
}

type CacheQuery struct {
	Name  string
	Key   string
	Value string
}
