package response

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type (
	Success struct {
		Success bool        `json:"success"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}

	notFound struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
		Data    any    `json:"data"`
		Error   any    `json:"error,omitempty"`
	}

	Failed struct {
		Success bool        `json:"success"`
		Message string      `json:"message"`
		Error   interface{} `json:"error"`
	}
	User struct {
		gorm.Model
		Email    string `json:"email"`
		Password string `json:"_"`
	}
)

func JsonSuccess(c *gin.Context, code int, data interface{}) {
	res := Success{
		Success: true,
		Message: "success",
		Data:    data,
	}

	c.JSON(code, res)
}

func JsonSuccessWithMessage(c *gin.Context, code int, message string, data any) {
	res := Success{
		Success: true,
		Message: message,
		Data:    data,
	}

	c.JSON(code, res)
}

func JsonFailed(c *gin.Context, code int, err error) {
	res := Failed{
		Success: false,
		Message: "failed",
		Error:   err,
	}

	c.JSON(code, res)
}

func JsonErrorWithMessage(c *gin.Context, code int, msg string, err error) {
	res := Failed{
		Success: false,
		Message: msg,
		Error:   err,
	}

	c.JSON(code, res)
}

func JsonErrorValidation(c *gin.Context, errs []error) {
	arrStringError := []string{}
	for _, err := range errs {
		arrStringError = append(arrStringError, err.Error())
	}

	res := Failed{
		Success: false,
		Message: "validation error",
		Error:   arrStringError,
	}

	c.JSON(400, res)
}

func JsonNotFound(c *gin.Context, err any) {
	res := notFound{
		Success: false,
		Message: "not found",
		Data:    []any{},
		Error:   err,
	}

	c.JSON(404, res)
}
