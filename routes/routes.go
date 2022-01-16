package routes

import (
	"gin-plus/app/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init() *gin.Engine {

	//创建一个默认的gin.Engine，使用了Logger()和Recovery()中间件
	r := gin.Default()
	//加入JWT中间件
	r.Use(middleware.JWT())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "success",
			"data": map[string]string{"ping": "pong"},
		})
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": http.StatusNotFound,
			"msg":  http.StatusText(http.StatusNotFound),
			"data": map[string]interface{}{},
		})
	})

	return r
}
