package services

import "github.com/himmel520/wb_L0/internal/repository"

type Redis interface {
	Set(id string, order []byte) error
	Get(id string) ([]byte, error)
}

type Order interface {
	Create(order []byte) (string, error)
	GetByID(id string) ([]byte, error)
}

type Service struct {
	Redis
	Order
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Redis: NewRedisService(repo.Redis),
		Order: NewOrderService(repo.Order),
	}
}
