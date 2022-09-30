package repository

import (
	"github.com/rudikurniawan99/go-api-4/src/model"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepository interface {
	Create(user *model.User) error
	FindByEmail(user *model.User, email string) error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(user *model.User) error {
	if err := r.db.Create(user).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepository) FindByEmail(user *model.User, email string) error {
	if err := r.db.Where("email = ?", email).First(user).Error; err != nil {
		return err
	}

	return nil
}
