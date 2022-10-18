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
	group.POST("register", d.RegisterHandler)
	group.GET("login", d.LoginHandler)
	group.GET("me", d.getMeHandler)
}

func (d *userDelivery) RegisterHandler(c *gin.Context) {
	req := &model.UserRequest{}
	c.Bind(req)

	if err := helper.UserValidator(req); err != nil {

		response.JsonErrorValidation(c, err.Error())
		return
	}

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

	token, err := helper.GenerateToken(int(user.ID))

	if err != nil {
		response.JsonErrorWithMessage(c, 400, "failed generate token", err)
	} else {

		c.Header("token", token)

		response.JsonSuccess(c, 200, gin.H{
			"token": token,
		})
	}

	// response.JsonSuccess(c, 200, user)

}

func (d *userDelivery) getMeHandler(c *gin.Context) {
	headerToken := c.GetHeader("token")
	id, errVal := helper.ValidateToken(headerToken)
	if errVal != nil {
		response.JsonErrorWithMessage(
			c,
			400,
			"err jwt",
			errVal,
		)
		return
	}

	user := &model.User{}
	ID, _ := strconv.Atoi(id)

	if err := d.u.FindById(user, ID); err != nil {
		response.JsonErrorWithMessage(
			c,
			404,
			"user not found",
			err,
		)
		return
	}

	response.JsonSuccess(
		c,
		200,
		gin.H{
			"user": user,
		},
	)

}
