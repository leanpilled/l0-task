package repository

import (
	"wb-0/entity"
	"wb-0/repository/cache"
)

type OrderRepo interface {
	GetOrders() ([]entity.Order, error)
	CreateOrder(order entity.Order) (entity.Order, error)
	FillCache(c *cache.Cache) error
}
