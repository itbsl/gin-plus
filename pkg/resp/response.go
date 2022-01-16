package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": CodeSuccess,
		"msg":  codeText[CodeSuccess],
		"data": data,
	})
}

func Fail(c *gin.Context, code int) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  codeText[code],
	})
}

func FailWithMsg(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

const (
	CodeSuccess      = 0
	CodeServerError  = 500
	CodeInvalidParam = 1000_000
)

var codeText = map[int]string{
	CodeSuccess:      "success",
	CodeInvalidParam: "参数错误",
	CodeServerError:  "服务器错误",
}
