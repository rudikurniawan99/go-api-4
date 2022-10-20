package delivery

import (
	"strconv"

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
	group.GET("/:id", d.getUserHanlder)
	group.DELETE("/:id", d.deleteUserByIdHandler)
	group.PUT("/:id", d.updateUserHandler)
}

func (d *userDelivery) getAllUserHanlder(c *gin.Context) {
	users := &[]model.User{}

	if err := d.u.GetAllUser(users); err != nil {
		response.JsonFailed(c, 400, err)
		return
	}

	response.JsonSuccess(c, 200, users)
}

func (d *userDelivery) getUserHanlder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		response.JsonFailed(c, 500, err)
		return
	}

	user := &model.User{}

	if err := d.u.FindById(user, id); err != nil {
		response.JsonNotFound(c, err)
		return
	}

	response.JsonSuccess(c, 200, user)
}

func (d *userDelivery) deleteUserByIdHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.JsonFailed(c, 500, err)
		return
	}

	user := &model.User{}

	if err := d.u.FindById(user, id); err != nil {
		response.JsonNotFound(c, err)
		return
	}

	if err := d.u.DeleteById(user, int(user.ID)); err != nil {
		response.JsonFailed(c, 500, err)
		return
	}

	response.JsonSuccessWithMessage(c, 200, "success delete user", []any{})
}

func (d *userDelivery) updateUserHandler(c *gin.Context) {
	req := &model.UserRequest{}
	c.Bind(req)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.JsonFailed(c, 500, err)
		return
	}

	// if err := helper.UserValidator(req); err != nil {
	// 	response.JsonErrorValidation(c, err)
	// 	return
	// }

	user := &model.User{}

	if err := d.u.FindById(user, id); err != nil {
		response.JsonNotFound(c, err)
		return
	}

	user.Email = req.Email

	if err := d.u.UpdateUser(user); err != nil {
		response.JsonErrorWithMessage(c, 400, "failed to update", err)
		return
	}

	response.JsonSuccess(c, 200, user)

}
