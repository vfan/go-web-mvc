package main

import (
	"log"
	"mvc-demo/config"
	"mvc-demo/controllers"
	"mvc-demo/dao"
	"mvc-demo/models"
	"mvc-demo/routes"
	"mvc-demo/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 应用依赖
type AppDependencies struct {
	DB             *gorm.DB
	UserDAO        *dao.UserDAO
	UserService    *service.UserService
	userController *controllers.UserController // 缓存控制器实例
}

// GetUserController 获取用户控制器
func (d *AppDependencies) GetUserController() *controllers.UserController {
	if d.userController == nil {
		d.userController = controllers.NewUserController(d.UserService)
	}
	return d.userController
}

func main() {
	// 获取配置
	appConfig := config.GetConfig()

	// 初始化数据库连接
	models.InitDB()
	log.Println("数据库初始化完成")

	// 初始化依赖
	deps := initDependencies(models.DB)

	// 设置Gin模式
	gin.SetMode(appConfig.Mode)

	// 设置路由
	r := routes.SetupRouter(deps)

	// 日志输出
	log.Printf("服务器启动于 %s 端口，运行模式: %s\n", appConfig.ServerPort, appConfig.Mode)

	// 启动服务器
	err := r.Run(appConfig.ServerPort)
	if err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}

// 初始化依赖
func initDependencies(db *gorm.DB) *AppDependencies {
	// 初始化DAO
	userDAO := dao.NewUserDAO(db)

	// 初始化服务
	userService := service.NewUserService(userDAO)

	return &AppDependencies{
		DB:          db,
		UserDAO:     userDAO,
		UserService: userService,
	}
}
