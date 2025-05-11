package main

import (
	"log"
	"mvc-demo/config"
	"mvc-demo/models"
	"mvc-demo/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// 获取配置
	appConfig := config.GetConfig()

	// 初始化数据库连接
	models.InitDB()
	log.Println("数据库初始化完成")

	// 设置Gin模式
	gin.SetMode(appConfig.Mode)

	// 设置路由
	r := routes.SetupRouter()

	// 日志输出
	log.Printf("服务器启动于 %s 端口，运行模式: %s\n", appConfig.ServerPort, appConfig.Mode)

	// 启动服务器
	err := r.Run(appConfig.ServerPort)
	if err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
