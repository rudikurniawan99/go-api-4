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
	authGroup := s.Group("auth")
	userDelivery.Mount(authGroup)
}
