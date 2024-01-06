package redis

import (
	"context"

	"github.com/himmel520/wb_L0/internal/config"
	"github.com/redis/go-redis/v9"
)

func NewRedis(cfg config.Redis) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: cfg.Addr,
	})
	ctx := context.Background()

	if _, err := rdb.Ping(ctx).Result(); err != nil {
		return nil, err
	}

	return rdb, nil
}
