package main

import (
	"context"
	"flag"
	"fmt"
	"gin-plus/global"
	"gin-plus/pkg/setting"
	"gin-plus/routes"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	//1.读取配置文件
	//支持通过命令行参数 -f 指定配置文件路径
	var filePath = flag.String("f", "", "配置文件路径")
	flag.Parse()
	setting.Init(*filePath)

	//2.初始化路由
	fmt.Println(global.Config.Mode)
	gin.SetMode(global.Config.Mode)
	router := routes.Init()

	//启动服务
	server := http.Server{
		Addr:           ":" + strconv.Itoa(global.Config.ServerConfig.Port),
		Handler:        router,
		ReadTimeout:    global.Config.ServerConfig.ReadTimeout * time.Second,
		WriteTimeout:   global.Config.ServerConfig.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		//开启一个goroutine启动服务
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP服务启动失败：server.ListenAndServe() failed: %v\n", err)
		}
	}()
	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1)
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit //阻塞在此，当接收到上述两种信号时才会往下执行
	log.Printf("Shutdown Server ...")
	//创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown err: %v\n", err)
	}
	log.Println("Server Exited")
}
