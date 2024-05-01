package cache

import (
	"auth/internal/config"
	"auth/internal/db"
	"context"
	"github.com/redis/go-redis/v9"
)

type Cache struct {
	cfg   config.RedisConfig
	cache *redis.Client
}

// todo add close

func New(cfg config.RedisConfig) db.Cache {
	options := &redis.Options{
		Addr:     cfg.Address(),
		Password: cfg.Password(),
		DB:       cfg.DB(),
	}
	return &Cache{
		cache: redis.NewClient(options),
	}
}

func (c *Cache) Get(ctx context.Context, key string) (string, error) {
	val := c.cache.Get(ctx, key)
	value, err := val.Result()
	if err != nil {
		return "", err
	}
	return value, nil
}

func (c *Cache) Add(ctx context.Context, key string, value string) error {
	add := c.cache.Append(ctx, key, value)
	_, err := add.Result()
	if err != nil {
		return err
	}
	return nil
}

func (c *Cache) Close() error {
	err := c.cache.Close()
	if err != nil {
		return err
	}
	return nil
}
