package services

import "github.com/himmel520/wb_L0/internal/repository"

type OrderService struct {
	repo repository.Order
}

func NewOrderService(repo repository.Order) *OrderService {
	return &OrderService{
		repo: repo,
	}
}

func (s *OrderService) Create(order []byte) (string, error) {
	return s.repo.Create(order)
}

func (s *OrderService) GetByID(id string) ([]byte, error) {
	return s.repo.GetByID(id)
}
