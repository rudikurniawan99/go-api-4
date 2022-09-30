package usecase

import (
	"github.com/rudikurniawan99/go-api-4/src/model"
	"github.com/rudikurniawan99/go-api-4/src/repository"
)

type (
	productUsecase struct {
		r repository.ProductRepository
	}

	ProductUsecase interface {
		CreateProduct(product *model.Product) error
	}
)

func NewProductUsecase(r repository.ProductRepository) ProductUsecase {
	return &productUsecase{r}
}

func (u *productUsecase) CreateProduct(product *model.Product) error {
	if err := u.r.Create(product); err != nil {
		return err
	}
	return nil
}
