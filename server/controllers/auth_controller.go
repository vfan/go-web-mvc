package controllers

import (
	"mvc-demo/config"
	"mvc-demo/models/dto"
	"mvc-demo/service"
	"mvc-demo/utils"

	"github.com/gin-gonic/gin"
)

// AuthController 认证控制器
type AuthController struct {
	userService *service.UserService
}

// NewAuthController 创建认证控制器
func NewAuthController(userService *service.UserService) *AuthController {
	return &AuthController{
		userService: userService,
	}
}

// Login 用户登录
func (a *AuthController) Login(c *gin.Context) {
	var req dto.LoginRequest

	// 请求参数绑定
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ParamError(c, "无效的请求参数")
		return
	}

	// 验证用户是否存在且密码正确
	user, err := a.userService.Login(req.Email, req.Password)
	if err != nil {
		utils.Unauthorized(c, "用户名或密码错误")
		return
	}

	// 生成JWT令牌
	token, err := utils.GenerateToken(user)
	if err != nil {
		utils.InternalError(c, "生成令牌失败")
		return
	}

	// 构建响应
	jwtConfig := config.GetJWTConfig()
	expiresIn := int(jwtConfig.TokenExpiry.Seconds())

	// 返回响应
	utils.SuccessWithMsg(c, "登录成功", dto.LoginResponse{
		Token:     token,
		TokenType: "Bearer",
		ExpiresIn: expiresIn,
	})
}
