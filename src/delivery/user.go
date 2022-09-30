package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/rudikurniawan99/go-api-4/src/model"
	"github.com/rudikurniawan99/go-api-4/src/response"
	"github.com/rudikurniawan99/go-api-4/src/usecase"
	"golang.org/x/crypto/bcrypt"
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
	group.POST("/register", d.CreateUserHandler)
}

func (d *userDelivery) CreateUserHandler(c *gin.Context) {
	req := &model.UserRequest{}
	c.Bind(req)

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 10)

	user := &model.User{
		Email:    req.Email,
		Password: string(hashPassword),
	}

	if err := d.u.CreateUser(user); err != nil {
		c.JSON(400, response.ErrorResponse{
			Code:    400,
			Message: "failed to create user",
			Error:   err,
		})

		return
	}

	c.JSON(201, response.SuccessResponse{
		Code:    201,
		Message: "success to create user",
		Data:    user,
	})
}
