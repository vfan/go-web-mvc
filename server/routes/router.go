package routes

import (
	"mvc-demo/controllers"
	"mvc-demo/middleware"

	"github.com/gin-gonic/gin"
)

// AppDependencies 应用依赖接口
type AppDependencies interface {
	GetUserController() *controllers.UserController
}

// SetupRouter 配置所有路由
func SetupRouter(deps AppDependencies) *gin.Engine {
	r := gin.Default()

	// 使用中间件
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// API 路由组
	api := r.Group("/api")
	{
		// 创建控制器实例
		authController := &controllers.AuthController{}
		userController := deps.GetUserController()
		universityController := &controllers.UniversityController{}
		studentController := &controllers.StudentController{}

		// 认证相关路由（无需认证）
		auth := api.Group("/auth")
		{
			auth.POST("/login", authController.Login)

			// 临时添加创建用户接口
			auth.POST("/register", userController.Create)
		}

		// 需要认证的路由
		authorized := api.Group("/")
		authorized.Use(middleware.JWTAuth())
		{
			// 用户路由
			userGroup := authorized.Group("/users")
			{
				userGroup.GET("/:id", userController.Get)
				userGroup.GET("", userController.List)
			}

			// 大学路由
			universityGroup := authorized.Group("/universities")
			{
				universityGroup.GET("/:id", universityController.Get)
				universityGroup.GET("", universityController.List)
				universityGroup.GET("/all", universityController.All)
			}

			// 学生路由
			studentGroup := authorized.Group("/students")
			{
				studentGroup.GET("/:id", studentController.Get)
				studentGroup.GET("", studentController.List)
			}

			// 需要管理员权限的路由
			admin := authorized.Group("/admin")
			admin.Use(middleware.AdminAuth())
			{
				// 用户管理
				admin.POST("/users", userController.Create)
				admin.PUT("/users/:id", userController.Update)
				admin.DELETE("/users/:id", userController.Delete)

				// 大学管理
				admin.POST("/universities", universityController.Create)
				admin.PUT("/universities/:id", universityController.Update)
				admin.DELETE("/universities/:id", universityController.Delete)

				// 学生管理
				admin.POST("/students", studentController.Create)
				admin.PUT("/students/:id", studentController.Update)
				admin.DELETE("/students/:id", studentController.Delete)
			}
		}
	}

	return r
}
