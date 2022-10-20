package usecase

import (
	"github.com/rudikurniawan99/go-api-4/src/model"
	"github.com/rudikurniawan99/go-api-4/src/repository"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	r repository.UserRepository
}

type UserUsecase interface {
	CreateUser(user *model.User) error
	FindByEmail(user *model.User, email string) error
	FindById(user *model.User, id int) error
	ComparePassword(password, hash string) error
	GetAllUser(users *[]model.User) error
}

func NewUserUsecase(r repository.UserRepository) UserUsecase {
	return &userUsecase{r}
}

func (u *userUsecase) CreateUser(user *model.User) error {

	if err := u.r.Create(user); err != nil {
		return err
	}

	return nil

}

func (u *userUsecase) FindByEmail(user *model.User, email string) error {
	if err := u.r.FindByEmail(user, email); err != nil {
		return err
	}

	return nil
}

func (u *userUsecase) ComparePassword(password, hash string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return err
	}

	return nil
}

func (u *userUsecase) FindById(user *model.User, id int) error {
	if err := u.r.FIndById(user, id); err != nil {
		return err
	}

	return nil
}

func (u *userUsecase) GetAllUser(users *[]model.User) error {
	if err := u.r.GetAll(users); err != nil {
		return err
	}

	return nil
}
