package entities

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Email     string    `json:"email" gorm:"not null;unique;size:255;index:idx_users_email"`
	Username  *string   `json:"username" gorm:"unique;size:64;index:idx_users_username"`
	Password  string    `json:"password" gorm:"not null;size:255;column:password"`
	FirstName *string   `json:"first_name" gorm:"size:100"`
	LastName  *string   `json:"last_name" gorm:"size:100"`
	Phone     *string   `json:"phone" gorm:"size:32"`
	IsActive  bool      `json:"is_active" gorm:"not null;default:true"`
	CreatedAt time.Time `json:"created_at" gorm:"not null;default:now()"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null;default:now()"`
}

type GetUserResponse struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Username  *string   `json:"username"`
	FirstName *string   `json:"first_name"`
	LastName  *string   `json:"last_name"`
	Phone     *string   `json:"phone"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateUserRequest struct {
	Email     string  `json:"email" binding:"required,email"`
	Username  *string `json:"username" binding:"omitempty,min=3,max=64"`
	Password  string  `json:"password" binding:"required,min=6,max=128"`
	FirstName *string `json:"first_name" binding:"omitempty,max=100"`
	LastName  *string `json:"last_name" binding:"omitempty,max=100"`
	Phone     *string `json:"phone" binding:"omitempty,max=32"`
}

func (u *User) HashPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
