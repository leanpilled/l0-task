package entity

import "gorm.io/gorm"

type Delivery struct {
	gorm.Model `json:"-"`
	OrderId    int    `gorm:"column:order_id"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Zip        string `json:"zip"`
	City       string `json:"city"`
	Address    string `json:"address"`
	Region     string `json:"region"`
	Email      string `json:"email"`
}
