package repository

import (
	"github.com/rudikurniawan99/go-api-4/src/model"
	"gorm.io/gorm"
)

type (
	productRepository struct {
		db *gorm.DB
	}

	ProductRepository interface {
		Create(product *model.Product) error
	}
)

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

func (r *productRepository) Create(product *model.Product) error {
	return nil
}
