package delivery

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rudikurniawan99/go-api-4/src/helper"
	"github.com/rudikurniawan99/go-api-4/src/model"
	"github.com/rudikurniawan99/go-api-4/src/response"
	"github.com/rudikurniawan99/go-api-4/src/usecase"
	"golang.org/x/crypto/bcrypt"
)

type authDelivery struct {
	u usecase.UserUsecase
}

type AuthDelivery interface {
	Mount(group *gin.RouterGroup)
}

func NewAuthDelivery(u usecase.UserUsecase) AuthDelivery {
	return &authDelivery{u}
}

func (d *authDelivery) Mount(group *gin.RouterGroup) {
	group.POST("login", d.Login)
	group.POST("register", d.Register)
	group.GET("me", d.getMe)
}

func (d *authDelivery) Login(c *gin.Context) {
	req := &model.UserRequest{}
	c.Bind(req)

	if err := helper.UserValidator(req); err != nil {
		response.JsonErrorValidation(c, err)
		return
	}

	user := &model.User{}

	if err := d.u.FindByEmail(user, req.Email); err != nil {
		response.JsonErrorWithMessage(c, 404, "email not found", err)
		return
	}

	if err := d.u.ComparePassword(req.Password, user.Password); err != nil {
		response.JsonErrorWithMessage(c, 401, "password not match", err)
		return
	}

	token, err := helper.GenerateToken(int(user.ID))
	if err != nil {
		response.JsonFailed(c, 500, err)
		return
	}

	response.JsonSuccess(c, 200, gin.H{
		"token": token,
	})
}

func (d *authDelivery) Register(c *gin.Context) {
	req := &model.UserRequest{}
	c.Bind(req)

	if err := helper.UserValidator(req); err != nil {
		response.JsonErrorValidation(c, err)
		return
	}

	hasPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	user := &model.User{
		Email:    req.Email,
		Password: string(hasPassword),
	}

	if err := d.u.FindByEmail(user, req.Email); err == nil {
		response.JsonErrorWithMessage(c, 400, "email already exist", err)
		return
	}

	if err := d.u.CreateUser(user); err != nil {
		response.JsonFailed(c, 400, err)
		return
	}

	token, err := helper.GenerateToken(int(user.ID))

	if err != nil {
		response.JsonFailed(c, 500, err)
		return
	}

	response.JsonSuccess(c, 201, gin.H{
		"data":  user,
		"token": token,
	})
}

func (d *authDelivery) getMe(c *gin.Context) {
	token := c.GetHeader("token")
	id, err := helper.ValidateToken(token)

	if token == "" || err != nil {
		response.JsonErrorWithMessage(c, 401, "unauthorized", nil)
		return
	}

	user := &model.User{}
	userId, _ := strconv.Atoi(id)

	if err := d.u.FindById(user, userId); err != nil {
		response.JsonErrorWithMessage(c, 404, "user not found", err)
		return
	}

	response.JsonSuccess(c, 200, user)
}
