package middleware

import (
	"gin-plus/pkg/app"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		//客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URL中
		//优先从请求头中获取Token
		token := c.GetHeader("token")
		if token == "" { //token中获取不到，则尝试从请求体中获取Token
			token = c.DefaultPostForm("token", "")
			if token == "" { //请求头和请求体中都获取不到，从URL获取
				token = c.DefaultQuery("token", "")
				if token == "" {
					c.JSON(http.StatusUnauthorized, gin.H{
						"code": http.StatusUnauthorized,
						"msg":  http.StatusText(http.StatusUnauthorized),
					})
					c.Abort()
					return
				}
			}
		}
		claims, err := app.Parse(token)
		if err != nil {
			switch err.(*jwt.ValidationError).Errors {
			case jwt.ValidationErrorExpired:
				c.JSON(http.StatusOK, gin.H{
					"code": 2001,
					"msg":  "Token过期",
				})
				c.Abort()
				return
			case jwt.ValidationErrorIssuer:
				c.JSON(http.StatusOK, gin.H{
					"code": 2002,
					"msg":  "签发者错误",
				})
				c.Abort()
				return
			default:
				c.JSON(http.StatusOK, gin.H{
					"code": 2003,
					"msg":  "无效的Token",
				})
				c.Abort()
				return
			}
		}
		//将当前请求的用户ID保存到请求的上下文中，后续的处理函数可以通过c.Get('user_id')来获取当前请求的用户信息
		c.Set("user_id", claims.UserId)

		c.Next()
	}
}
