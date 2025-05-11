package routes

import (
	"mvc-demo/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter 配置所有路由
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 使用中间件
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// API 路由组
	api := r.Group("/api")
	{
		// 创建控制器实例
		userController := &controllers.UserController{}
		universityController := &controllers.UniversityController{}
		studentController := &controllers.StudentController{}

		// 用户路由
		userGroup := api.Group("/users")
		{
			userGroup.POST("", userController.Create)
			userGroup.GET("/:id", userController.Get)
			userGroup.PUT("/:id", userController.Update)
			userGroup.DELETE("/:id", userController.Delete)
			userGroup.GET("", userController.List)
		}

		// 大学路由
		universityGroup := api.Group("/universities")
		{
			universityGroup.POST("", universityController.Create)
			universityGroup.GET("/:id", universityController.Get)
			universityGroup.PUT("/:id", universityController.Update)
			universityGroup.DELETE("/:id", universityController.Delete)
			universityGroup.GET("", universityController.List)
			universityGroup.GET("/all", universityController.All)
		}

		// 学生路由
		studentGroup := api.Group("/students")
		{
			studentGroup.POST("", studentController.Create)
			studentGroup.GET("/:id", studentController.Get)
			studentGroup.PUT("/:id", studentController.Update)
			studentGroup.DELETE("/:id", studentController.Delete)
			studentGroup.GET("", studentController.List)
		}
	}

	return r
}
