package main

import (
	"flag"
	"gin-plus/global"
	"gin-plus/pkg/setting"
	"gin-plus/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	//1.读取配置文件
	//支持通过命令行参数 -f 指定配置文件路径
	var filePath = flag.String("f", "", "配置文件路径")
	flag.Parse()
	setting.Init(*filePath)

	//2.初始化路由
	gin.SetMode(global.Config.Mode)
	router := routes.Init()

	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("项目启动失败:%v\n", err)
	}
}
