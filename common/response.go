package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Ok(c *gin.Context) {
	c.JSON(http.StatusOK, &Response{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    nil,
	})
}

func OkWithMessage(c *gin.Context, message string) {
	c.JSON(http.StatusOK, &Response{
		Code:    http.StatusOK,
		Message: message,
		Data:    nil,
	})
}

func OkWithData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &Response{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    data,
	})
}

func Failed(c *gin.Context, code int) {
	c.JSON(code, &Response{
		Code:    code,
		Message: Message(code),
		Data:    nil,
	})
}

func FailedWithMessage(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, &Response{
		Code:    http.StatusInternalServerError,
		Message: message,
		Data:    nil,
	})
}
