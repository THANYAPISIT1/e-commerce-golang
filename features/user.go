package services

import (
	"e-commerce/entities"

	"gorm.io/gorm"
)

var db *gorm.DB

func GetUsers() ([]entities.User, error) {
	var users []entities.User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
