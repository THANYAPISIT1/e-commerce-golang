package domain

import (
	"e-commerce/entities"
)

type UserUsecase interface {
	GetUser(id uint32) (entities.GetUserResponse, error)
	CreateUser(req entities.CreateUserRequest) error
}

type UserRepository interface {
	CreateUser(req entities.CreateUserRequest) error
	GetUser(id uint32) (entities.GetUserResponse, error)
}
