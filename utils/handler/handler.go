package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/liyinda/google-authenticator/utils/errno"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)

	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func NotFound(c *gin.Context) {
	SendResponse(c, errno.ErrNotFound, nil)
	/*
		code, message := errno.DecodeErr(errno.ErrNotFound)
		c.JSON(404, Response{
			Code:    code,
			Message: message,
			Data:    nil,
		})
	*/
}
