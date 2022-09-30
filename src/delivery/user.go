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
	group.POST("register", d.CreateUserHandler)
	group.GET("login", d.LoginHandler)
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
		response.JsonFailed(c, 400, err)
		return
	}

	response.JsonSuccess(c, 201, user)
}

func (d *userDelivery) LoginHandler(c *gin.Context) {
	req := &model.UserRequest{}
	err := c.BindJSON(req)

	if err != nil {
		response.JsonFailed(c, 400, err)
		return
	}

	user := &model.User{}

	if err := d.u.FindByEmail(user, req.Email); err != nil {
		response.JsonErrorWithMessage(c, 404, "email not found", err)
		return
	}

	if err := d.u.ComparePassword(req.Password, user.Password); err != nil {
		response.JsonErrorWithMessage(c, 400, "password not match", err)
		return
	}

	response.JsonSuccess(c, 200, user)
}
