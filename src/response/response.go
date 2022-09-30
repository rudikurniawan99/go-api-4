package response

type (
	SuccessResponse struct {
		Code    int         `json:"code"`
		Message string      `json:"messsage"`
		Data    interface{} `json:"data"`
	}

	ErrorResponse struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Error   interface{} `json:"error"`
	}
)
