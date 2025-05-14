package main

import (
	"log"
	"mvc-demo/config"
	"mvc-demo/controllers"
	"mvc-demo/dao"
	"mvc-demo/db"
	"mvc-demo/routes"
	"mvc-demo/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 应用依赖
type AppDependencies struct {
	DB                   *gorm.DB
	UserDAO              *dao.UserDAO
	UniversityDAO        *dao.UniversityDAO
	StudentDAO           *dao.StudentDAO
	UserService          *service.UserService
	UniversityService    *service.UniversityService
	StudentService       *service.StudentService
	userController       *controllers.UserController
	authController       *controllers.AuthController
	universityController *controllers.UniversityController
	studentController    *controllers.StudentController
}

// GetUserController 获取用户控制器
func (d *AppDependencies) GetUserController() *controllers.UserController {
	if d.userController == nil {
		d.userController = controllers.NewUserController(d.UserService)
	}
	return d.userController
}

// GetAuthController 获取认证控制器
func (d *AppDependencies) GetAuthController() *controllers.AuthController {
	if d.authController == nil {
		d.authController = controllers.NewAuthController(d.UserService)
	}
	return d.authController
}

// GetUniversityController 获取大学控制器
func (d *AppDependencies) GetUniversityController() *controllers.UniversityController {
	if d.universityController == nil {
		d.universityController = controllers.NewUniversityController(d.UniversityService)
	}
	return d.universityController
}

// GetStudentController 获取学生控制器
func (d *AppDependencies) GetStudentController() *controllers.StudentController {
	if d.studentController == nil {
		d.studentController = controllers.NewStudentController(d.StudentService, d.UniversityService)
	}
	return d.studentController
}

func main() {
	// 获取配置
	appConfig := config.GetConfig()

	// 初始化数据库连接
	db.InitDB()
	log.Println("数据库初始化完成")

	// 初始化依赖
	deps := initDependencies(db.DB)

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
	universityDAO := dao.NewUniversityDAO(db)
	studentDAO := dao.NewStudentDAO(db)

	// 初始化服务
	userService := service.NewUserService(userDAO)
	universityService := service.NewUniversityService(universityDAO)
	studentService := service.NewStudentService(studentDAO)

	return &AppDependencies{
		DB:                db,
		UserDAO:           userDAO,
		UniversityDAO:     universityDAO,
		StudentDAO:        studentDAO,
		UserService:       userService,
		UniversityService: universityService,
		StudentService:    studentService,
	}
}
