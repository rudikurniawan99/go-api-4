package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/rudikurniawan99/go-api-4/src/usecase"
)

type (
	productDelivery struct {
		u usecase.ProductUsecase
	}

	ProductDelivery interface {
		Mount(group *gin.RouterGroup)
	}
)

func NewProductDelivery(u usecase.ProductUsecase) ProductDelivery {
	return &productDelivery{u}
}

func (d *productDelivery) Mount(group *gin.RouterGroup) {
	group.POST("/", d.createHandler)
}

func (d *productDelivery) createHandler(c *gin.Context) {

}
