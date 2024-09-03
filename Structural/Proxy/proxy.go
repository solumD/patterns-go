package proxy

import (
	"fmt"
)

// интерфейс сервиса
type Service interface {
	GetAllFruits(string) ([]string, error)
}

// реальный сервис
type GoodsService struct {
}

// возвращает товары
func (g GoodsService) GetAllFruits(key string) ([]string, error) {
	return []string{"apples", "oranges", "bananas"}, nil
}

// заместитель реального сервиса
type GoodsServiceProxy struct {
	realService Service
}

// перехватывает вызов к реальному объекту, и, если ключ не верен, возращает ошибку,
// иначе вызывает метод реального сервиса и возвращает его результат
func (g GoodsServiceProxy) GetAllFruits(key string) ([]string, error) {
	if key != "go is awesome" {
		return nil, fmt.Errorf("invalid key. access denied")
	}
	return g.realService.GetAllFruits(key)
}
