package database

import (
	"Katara/internal/config"
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

func RedisLoad(cfg *config.Config) (*redis.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.REDIS_URI,
		Password: cfg.REDIS_PASSWORD,
		DB:       int(cfg.REDIS_DB),
	})

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return rdb, nil
}
