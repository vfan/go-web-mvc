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

	// 设置HttpOnly Cookie
	c.SetCookie(
		"token",                         // cookie名称
		token,                           // cookie值
		expiresIn,                       // 过期时间（秒）
		"/",                             // 路径
		"",                              // 域名（空表示当前域名）
		c.Request.URL.Scheme == "https", // 仅在HTTPS连接时启用安全标志
		true,                            // HttpOnly, 禁止JavaScript访问
	)

	// 返回响应（不包含token，但包含其他信息）
	utils.SuccessWithMsg(c, "登录成功", dto.LoginResponse{
		TokenType: "Bearer",
		ExpiresIn: expiresIn,
	})
}

// Logout 用户登出
func (a *AuthController) Logout(c *gin.Context) {
	// 清除token Cookie
	c.SetCookie(
		"token",                         // cookie名称
		"",                              // cookie值设置为空
		-1,                              // 过期时间设为负数，立即过期
		"/",                             // 路径
		"",                              // 域名（空表示当前域名）
		c.Request.URL.Scheme == "https", // 仅在HTTPS连接时启用安全标志
		true,                            // HttpOnly，禁止JavaScript访问
	)

	utils.Success(c, nil)
}

// Me 获取当前登录用户信息
func (a *AuthController) Me(c *gin.Context) {
	// 从上下文中获取用户ID
	userID, exists := c.Get("user_id")
	if !exists {
		utils.Unauthorized(c, "未登录")
		return
	}

	// 获取用户信息，需要将userID转换为int64类型
	user, err := a.userService.GetUserByID(int64(userID.(uint)))
	if err != nil {
		utils.InternalError(c, "获取用户信息失败")
		return
	}

	// 返回用户信息
	utils.Success(c, map[string]interface{}{
		"id":       user.ID,
		"email":    user.Email,
		"username": user.Username,
		"role":     user.Role,
	})
}
