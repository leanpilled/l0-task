package cache

import "wb-0/entity"

type Cache struct {
	cache map[int]entity.Order
}

func NewCache() *Cache {
	return &Cache{cache: make(map[int]entity.Order)}
}

func (c *Cache) GetByID(id int) (entity.Order, bool) {
	val, ok := c.cache[id]
	return val, ok
}

func (c *Cache) Set(order entity.Order) {
	c.cache[order.Id] = order
}
