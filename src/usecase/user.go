package usecase

import (
	"github.com/rudikurniawan99/go-api-4/src/model"
	"github.com/rudikurniawan99/go-api-4/src/repository"
)

type userUsecase struct {
	r repository.UserRepository
}

type UserUsecase interface {
	CreateUser(user *model.User) error
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
