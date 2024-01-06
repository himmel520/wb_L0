package services

import "github.com/himmel520/wb_L0/internal/repository"

type RedisService struct {
	repo repository.Redis
}

func NewRedisService(repo repository.Redis) *RedisService {
	return &RedisService{
		repo: repo,
	}
}

func (s *RedisService) Set(id string, order []byte) error {
	return s.repo.Set(id, order)
}

func (s *RedisService) Get(id string) ([]byte, error) {
	return s.repo.Get(id)
}
