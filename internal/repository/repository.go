package repository

import (
	"database/sql"

	"github.com/himmel520/wb_L0/internal/repository/postgres"
	redisRepo "github.com/himmel520/wb_L0/internal/repository/redis"
	"github.com/redis/go-redis/v9"
)

type Redis interface {
	Set(id string, order []byte) error
	Get(id string) ([]byte, error)
}

type Order interface {
	Create(order []byte) (string, error)
	GetByID(id string) ([]byte, error)
}

type Repository struct {
	Order
	Redis
}

func NewRepository(db *sql.DB, rdb *redis.Client) *Repository {
	return &Repository{
		Order: postgres.NewOrderPg(db),
		Redis: redisRepo.NewOrderRedis(rdb),
	}
}
