package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int64       `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func Rep(c *gin.Context, code int64, data interface{}, msg string) {
	httpcode := http.StatusOK
	if code > 2000 {
		httpcode = http.StatusRequestEntityTooLarge
	}
	c.JSON(httpcode, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
	return
}
