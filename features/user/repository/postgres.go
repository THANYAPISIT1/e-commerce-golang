package repository

import (
	"e-commerce/domain"
	"e-commerce/entities"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(req entities.CreateUserRequest) error {
	if err := r.db.Table("users").Create(&req).Error; err != nil {
		return errors.Wrap(err, "[UserRepository.CreateUser]: failed to create user")
	}

	return nil
}

func (r *userRepository) GetUser(id uint32) (entities.GetUserResponse, error) {
	var user entities.GetUserResponse
	if err := r.db.Table("users").Select("users.*, companies.name as company_name").Joins("JOIN companies ON users.company_id = companies.id").
		First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, nil
		}

		return user, errors.Wrap(err, "[UserRepository.GetUser]: failed to get user")
	}

	return user, nil
}
