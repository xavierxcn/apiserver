package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xavierxcn/apiserver/internal/serve/pkg/errno"
)

// Response resp
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type defaultResponse struct {
}

// SendResponse response
func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)
	if data == nil {
		data = defaultResponse{}
	}

	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
