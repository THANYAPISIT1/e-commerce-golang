package usecase

import (
	"e-commerce/domain"
	"e-commerce/entities"
	"strings"

	"github.com/pkg/errors"
)

type userUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(userRepo domain.UserRepository) domain.UserUsecase {
	return &userUsecase{userRepo: userRepo}
}

func (u *userUsecase) GetUser(id uint32) (entities.GetUserResponse, error) {
	user, err := u.userRepo.GetUser(id)
	if err != nil {
		return entities.GetUserResponse{}, errors.Wrap(err, "[UserUsecase.GetUser]: failed to get user")
	}

	return user, nil
}

func (u *userUsecase) CreateUser(req entities.CreateUserRequest) error {
	tempUser := entities.User{}
	if err := tempUser.HashPassword(req.Password); err != nil {
		return errors.Wrap(err, "[UserUsecase.CreateUser]: failed to hash password")
	}

	req.Password = tempUser.Password

	if err := u.userRepo.CreateUser(req); err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			if strings.Contains(err.Error(), "uni_users_email") {
				return errors.New("user with this email already exists")
			}
			if strings.Contains(err.Error(), "uni_users_username") {
				return errors.New("user with this username already exists")
			}
			return errors.New("user already exists")
		}
		return errors.Wrap(err, "[UserUsecase.CreateUser]: failed to create user")
	}
	return nil
}
