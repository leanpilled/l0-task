package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPGConn(DATABASE_URL string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(DATABASE_URL), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to setup a new connection")
		return nil, err
	}
	return db, nil
}
