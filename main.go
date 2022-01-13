package main

import (
	"gin-plus/routes"
	"log"
)

func main() {
	//初始化路由
	router := routes.Init()

	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("项目启动失败:%v\n", err)
	}
}
