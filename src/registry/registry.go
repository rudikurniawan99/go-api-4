package registry

import (
	"github.com/gin-gonic/gin"
	"github.com/rudikurniawan99/go-api-4/src/delivery"
	"github.com/rudikurniawan99/go-api-4/src/repository"
	"github.com/rudikurniawan99/go-api-4/src/usecase"
	"gorm.io/gorm"
)

func Registry(s *gin.Engine, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userDelivery := delivery.NewUserDelivery(userUsecase)
	authDelivery := delivery.NewAuthDelivery(userUsecase)
	authGroup := s.Group("auth")
	userGroup := s.Group("users")
	userDelivery.Mount(userGroup)
	authDelivery.Mount(authGroup)

	productRepository := repository.NewProductRepository(db)
	productUsecase := usecase.NewProductUsecase(productRepository)
	productDelivery := delivery.NewProductDelivery(productUsecase)
	productGroup := s.Group("products")
	productDelivery.Mount(productGroup)
}
