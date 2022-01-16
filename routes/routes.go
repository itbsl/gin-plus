package routes

import (
	"gin-plus/app/middleware"
	"gin-plus/pkg/resp"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init() *gin.Engine {

	//创建一个默认的gin.Engine，使用了Logger()和Recovery()中间件
	r := gin.Default()
	//加入JWT中间件，翻译中间件
	r.Use(middleware.JWT(), middleware.Translations())

	r.GET("/ping", func(c *gin.Context) {
		resp.Success(c, map[string]string{"ping": "ping"})
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": http.StatusNotFound,
			"msg":  http.StatusText(http.StatusNotFound),
		})
	})

	return r
}
