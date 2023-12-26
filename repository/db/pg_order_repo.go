package db

import (
	"wb-0/entity"
	"wb-0/repository/cache"

	"gorm.io/gorm"
)

type PGOrderRepo struct {
	db *gorm.DB
}

func NewPGOrderRepo(db *gorm.DB) *PGOrderRepo {
	return &PGOrderRepo{db: db}
}

func (repo *PGOrderRepo) GetOrders() ([]entity.Order, error) {
	var order []entity.Order
	err := repo.db.Preload("Delivery").Preload("Payment").Preload("Items").Find(&order).Error
	return order, err
}

func (repo *PGOrderRepo) CreateOrder(order entity.Order) (entity.Order, error) {
	err := repo.db.Create(&order).Error
	if err != nil {
		return order, err
	}
	return order, nil
}

func (repo *PGOrderRepo) FillCache(c *cache.Cache) error {
	orders, err := repo.GetOrders()
	if err != nil {
		return err
	}
	for i := range orders {
		c.Set(orders[i])
	}
	return nil
}
