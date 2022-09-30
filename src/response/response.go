package response

import "github.com/gin-gonic/gin"

type (
	Success struct {
		Success bool        `json:"success"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}

	Failed struct {
		Success bool        `json:"success"`
		Message string      `json:"message"`
		Error   interface{} `json:"error"`
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
