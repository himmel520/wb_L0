package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type OrderRedis struct {
	rdb *redis.Client
	ctx context.Context
}

func NewOrderRedis(rdb *redis.Client) *OrderRedis {
	return &OrderRedis{
		rdb: rdb,
		ctx: context.Background(),
	}
}

func (r *OrderRedis) Set(id string, order []byte) error {
	return r.rdb.Set(
		r.ctx,
		fmt.Sprintf("orders:%v", id),
		order,
		time.Hour*24*7).Err()
}

func (r *OrderRedis) Get(id string) ([]byte, error) {
	order, err := r.rdb.Get(
		r.ctx,
		fmt.Sprintf("orders:%v", id),
	).Bytes()
	if err != nil {
		return nil, err
	}
	return order, nil
}
