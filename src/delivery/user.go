package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/rudikurniawan99/go-api-4/src/model"
	"github.com/rudikurniawan99/go-api-4/src/response"
	"github.com/rudikurniawan99/go-api-4/src/usecase"
)

type userDelivery struct {
	u usecase.UserUsecase
}

type UserDelivery interface {
	Mount(group *gin.RouterGroup)
}

func NewUserDelivery(u usecase.UserUsecase) UserDelivery {
	return &userDelivery{u}
}

func (d *userDelivery) Mount(group *gin.RouterGroup) {
	group.GET("/", d.getAllUserHanlder)
}

func (d *userDelivery) getAllUserHanlder(c *gin.Context) {
	users := &[]model.User{}

	if err := d.u.GetAllUser(users); err != nil {
		response.JsonFailed(c, 400, err)
		return
	}

	response.JsonSuccess(c, 200, users)
}
