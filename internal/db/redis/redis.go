package cache

import (
	"auth/internal/config"
	"auth/internal/db"
	"auth/internal/logger"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type Cache struct {
	cfg   config.RedisConfig
	cache *redis.Client
	log   logger.Logger
}

// todo add close

func New(cfg config.RedisConfig, log logger.Logger) *Cache {
	const op = "cache.New"

	options := &redis.Options{
		Addr:     cfg.Address(),
		Password: cfg.Password(),
		DB:       cfg.DB(),
	}

	c := &Cache{
		cache: redis.NewClient(options),
		log:   log,
	}

	status := c.cache.Ping(context.Background())

	if s, err := status.Result(); s != "PONG" {
		log.FatalOp(op, err)
	}
	log.Info("start redis on " + cfg.Address())

	return c
}

func (c *Cache) logQuery(q db.CacheQuery) {
	l := fmt.Sprintf("noSql: %s, key: %s, value: %s", q.Name, q.Key, q.Value)

	c.log.Info(l)
}

func (c *Cache) Get(ctx context.Context, query db.CacheQuery) (string, error) {
	c.logQuery(query)

	val := c.cache.Get(ctx, query.Key)
	value, err := val.Result()
	if err != nil {
		return "", err
	}
	return value, nil
}

func (c *Cache) Add(ctx context.Context, query db.CacheQuery) error {
	c.logQuery(query)

	add := c.cache.Append(ctx, query.Key, query.Value)
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
